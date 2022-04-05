package gateway

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	"github.com/runatlantis/atlantis/server/core/config/valid"
	"github.com/runatlantis/atlantis/server/events"
	"github.com/runatlantis/atlantis/server/events/command"
	"github.com/runatlantis/atlantis/server/events/metrics"
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/events/vcs"
	"github.com/runatlantis/atlantis/server/logging"
	"github.com/runatlantis/atlantis/server/recovery"
	"github.com/uber-go/tally"
)

// AutoplanValidator handles setting up repo cloning and checking to verify of any terraform files have changed
type AutoplanValidator struct {
	Scope                         tally.Scope
	VCSClient                     vcs.Client
	PreWorkflowHooksCommandRunner events.PreWorkflowHooksCommandRunner
	Drainer                       *events.Drainer
	GlobalCfg                     valid.GlobalCfg
	CommitStatusUpdater           events.CommitStatusUpdater
	PrjCmdBuilder                 events.ProjectPlanCommandBuilder
	PullUpdater                   *events.PullUpdater
	WorkingDir                    events.WorkingDir
	WorkingDirLocker              events.WorkingDirLocker
}

const DefaultWorkspace = "default"

func (r *AutoplanValidator) createLogger(logger logging.SimpleLogging, repoName string, pullNum int) logging.SimpleLogging {
	return logger.With(
		"repository", repoName,
		"pull-num", strconv.Itoa(pullNum),
	)
}

func (r *AutoplanValidator) isValid(logger logging.SimpleLogging, baseRepo models.Repo, headRepo models.Repo, pull models.PullRequest, user models.User) (bool, error) {
	if opStarted := r.Drainer.StartOp(); !opStarted {
		return false, errors.New("atlantis is shutting down, cannot process current event")
	}
	defer r.Drainer.OpDone()

	log := r.createLogger(logger, baseRepo.FullName, pull.Num)
	defer r.logPanics(log)

	ctx := &command.Context{
		User:     user,
		Log:      log,
		Scope:    r.Scope,
		Pull:     pull,
		HeadRepo: headRepo,
		Trigger:  command.AutoTrigger,
	}
	if !r.validateCtxAndComment(ctx) {
		return false, errors.New("invalid command context")
	}
	err := r.PreWorkflowHooksCommandRunner.RunPreHooks(context.TODO(), ctx)
	if err != nil {
		ctx.Log.Errorf("Error running pre-workflow hooks %s. Proceeding with %s command.", err, command.Plan)
	}

	projectCmds, err := r.PrjCmdBuilder.BuildAutoplanCommands(ctx)
	if err != nil {
		if statusErr := r.CommitStatusUpdater.UpdateCombined(context.TODO(), baseRepo, pull, models.FailedCommitStatus, command.Plan); statusErr != nil {
			ctx.Log.Warnf("unable to update commit status: %w", statusErr)
		}
		// If error happened after clone was made, we should clean it up here too
		unlockFn, lockErr := r.WorkingDirLocker.TryLock(baseRepo.FullName, pull.Num, DefaultWorkspace)
		if lockErr != nil {
			ctx.Log.Warnf("workspace was locked")
			return false, errors.Wrap(err, lockErr.Error())
		}
		defer unlockFn()
		if cloneErr := r.WorkingDir.Delete(baseRepo, pull); cloneErr != nil {
			ctx.Log.With("err", cloneErr).Warnf("unable to delete clone after autoplan failed")
		}
		r.PullUpdater.UpdatePull(ctx, events.AutoplanCommand{}, command.Result{Error: err})
		return false, errors.Wrap(err, "Failed building autoplan commands")
	}
	unlockFn, err := r.WorkingDirLocker.TryLock(baseRepo.FullName, pull.Num, DefaultWorkspace)
	if err != nil {
		ctx.Log.Warnf("workspace was locked")
		return false, err
	}
	defer unlockFn()
	// Delete repo clone generated to validate plan
	if err := r.WorkingDir.Delete(baseRepo, pull); err != nil {
		return false, errors.Wrap(err, "Failed deleting cloned repo")
	}
	if len(projectCmds) == 0 {
		ctx.Log.Infof("no modified projects have been found")
		for _, cmd := range []command.Name{command.Plan, command.Apply, command.PolicyCheck} {
			if err := r.CommitStatusUpdater.UpdateCombinedCount(context.TODO(), baseRepo, pull, models.SuccessCommitStatus, cmd, 0, 0); err != nil {
				ctx.Log.Warnf("unable to update commit status: %s", err)
			}
		}
		return false, nil
	}
	return true, nil
}

func (r *AutoplanValidator) InstrumentedIsValid(logger logging.SimpleLogging, baseRepo models.Repo, headRepo models.Repo, pull models.PullRequest, user models.User) bool {
	timer := r.Scope.Timer(metrics.ExecutionTimeMetric).Start()
	defer timer.Stop()
	isValid, err := r.isValid(logger, baseRepo, headRepo, pull, user)

	log := r.createLogger(logger, baseRepo.FullName, pull.Num)

	if err != nil {
		log.Errorf(err.Error())
		r.Scope.Counter(metrics.ExecutionErrorMetric).Inc(1)
		return false
	}
	if !isValid {
		r.Scope.Counter(metrics.ExecutionFailureMetric).Inc(1)
		return false
	}
	r.Scope.Counter(metrics.ExecutionSuccessMetric).Inc(1)
	return true
}

func (r *AutoplanValidator) logPanics(logger logging.SimpleLogging) {
	if err := recover(); err != nil {
		stack := recovery.Stack(3)
		logger.Errorf("PANIC: %s\n%s", err, stack)
	}
}

func (r *AutoplanValidator) validateCtxAndComment(ctx *command.Context) bool {
	if ctx.HeadRepo.Owner != ctx.Pull.BaseRepo.Owner {
		ctx.Log.Infof("command was run on a fork pull request which is disallowed")
		return false
	}

	if ctx.Pull.State != models.OpenPullState {
		ctx.Log.Infof("command was run on closed pull request")
		return false
	}

	repo := r.GlobalCfg.MatchingRepo(ctx.Pull.BaseRepo.ID())
	if !repo.BranchMatches(ctx.Pull.BaseBranch) {
		ctx.Log.Infof("command was run on a pull request which doesn't match base branches")
		// just ignore it to allow us to use any git workflows without malicious intentions.
		return false
	}
	return true
}