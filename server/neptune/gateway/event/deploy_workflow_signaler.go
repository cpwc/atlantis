package event

import (
	"context"
	"fmt"

	"github.com/runatlantis/atlantis/server/core/config/valid"
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/neptune/workflows"
	"github.com/runatlantis/atlantis/server/vcs"
	"go.temporal.io/sdk/client"
)

type signaler interface {
	SignalWithStartWorkflow(ctx context.Context, workflowID string, signalName string, signalArg interface{},
		options client.StartWorkflowOptions, workflow interface{}, workflowArgs ...interface{}) (client.WorkflowRun, error)
	SignalWorkflow(ctx context.Context, workflowID string, runID string, signalName string, arg interface{}) error
}

const (
	Deprecated = "deprecated"
	Destroy    = "-destroy"
)

type DeployWorkflowSignaler struct {
	TemporalClient signaler
}

func (d *DeployWorkflowSignaler) SignalWithStartWorkflow(
	ctx context.Context,
	rootCfg *valid.MergedProjectCfg,
	repo models.Repo,
	revision string,
	installationToken int64,
	ref vcs.Ref,
	sender models.User,
	trigger workflows.Trigger) (client.WorkflowRun, error) {

	options := client.StartWorkflowOptions{TaskQueue: workflows.DeployTaskQueue}

	var tfVersion string
	if rootCfg.TerraformVersion != nil {
		tfVersion = rootCfg.TerraformVersion.String()
	}

	run, err := d.TemporalClient.SignalWithStartWorkflow(
		ctx,
		buildDeployWorkflowID(repo.FullName, rootCfg.Name),
		workflows.DeployNewRevisionSignalID,
		workflows.DeployNewRevisionSignalRequest{
			Revision: revision,
			InitiatingUser: workflows.User{
				Name: sender.Username,
			},
			Root: workflows.Root{
				Name: rootCfg.Name,
				Plan: workflows.Job{
					Steps: d.generateSteps(rootCfg.DeploymentWorkflow.Plan.Steps),
				},
				Apply: workflows.Job{
					Steps: d.generateSteps(rootCfg.DeploymentWorkflow.Apply.Steps),
				},
				RepoRelPath: rootCfg.RepoRelDir,
				TfVersion:   tfVersion,
				PlanMode:    d.generatePlanMode(rootCfg),
				Trigger:     trigger,
			},
			Repo: workflows.Repo{
				URL:      repo.CloneURL,
				FullName: repo.FullName,
				Name:     repo.Name,
				Owner:    repo.Owner,
				Credentials: workflows.AppCredentials{
					InstallationToken: installationToken,
				},
				Ref: workflows.Ref{
					Name: ref.Name,
					Type: string(ref.Type),
				},
			},
			Tags: rootCfg.Tags,
		},
		options,
		workflows.Deploy,
		workflows.DeployRequest{},
	)
	return run, err
}

func buildDeployWorkflowID(repoName string, rootName string) string {
	return fmt.Sprintf("%s||%s", repoName, rootName)
}

func (d *DeployWorkflowSignaler) generateSteps(steps []valid.Step) []workflows.Step {
	// NOTE: for deployment workflows, we won't support command level user requests for log level output verbosity
	var workflowSteps []workflows.Step
	for _, step := range steps {
		workflowSteps = append(workflowSteps, workflows.Step{
			StepName:    step.StepName,
			ExtraArgs:   step.ExtraArgs,
			RunCommand:  step.RunCommand,
			EnvVarName:  step.EnvVarName,
			EnvVarValue: step.EnvVarValue,
		})
	}
	return workflowSteps
}

func (d *DeployWorkflowSignaler) generatePlanMode(cfg *valid.MergedProjectCfg) workflows.PlanMode {
	t, ok := cfg.Tags[Deprecated]
	if ok && t == Destroy {
		return workflows.DestroyPlanMode
	}

	return workflows.NormalPlanMode
}