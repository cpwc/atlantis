package terraform

import (
	"context"

	"github.com/runatlantis/atlantis/server/neptune/workflows/internal/activities"
	"github.com/runatlantis/atlantis/server/neptune/workflows/internal/config/logger"
	"github.com/runatlantis/atlantis/server/neptune/workflows/internal/github"
	"github.com/runatlantis/atlantis/server/neptune/workflows/internal/github/markdown"
	"github.com/runatlantis/atlantis/server/neptune/workflows/internal/terraform/state"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type receiverActivities interface {
	UpdateCheckRun(ctx context.Context, request activities.UpdateCheckRunRequest) (activities.UpdateCheckRunResponse, error)
}

type StateReceiver struct {
	Repo     github.Repo
	Activity receiverActivities
}

func (n *StateReceiver) Receive(ctx workflow.Context, c workflow.ReceiveChannel, deploymentInfo DeploymentInfo) {
	var workflowState *state.Workflow
	c.Receive(ctx, &workflowState)

	// TODO: if we never created a check run, there was likely some issue, we should attempt to create it again.
	if deploymentInfo.CheckRunID == 0 {
		logger.Error(ctx, "check run id is 0, skipping update of check run")
		return
	}

	// this shouldn't be possible
	if workflowState.Plan == nil {
		logger.Error(ctx, "Plan job state is nil, This is likely a bug. Unable to update checks")
		return
	}

	summary := markdown.RenderWorkflowStateTmpl(workflowState)
	checkRunState := determineCheckRunState(workflowState)

	request := activities.UpdateCheckRunRequest{
		Title:   "atlantis/deploy",
		State:   checkRunState,
		Repo:    n.Repo,
		ID:      deploymentInfo.CheckRunID,
		Summary: summary,
	}

	if workflowState.Plan.Status == state.SuccessJobStatus && workflowState.Apply == nil {
		request.Actions = []github.CheckRunAction{
			github.CreatePlanReviewAction(github.Approve),
			github.CreatePlanReviewAction(github.Reject),
		}
	}

	// cap our retries for non-terminal states to allow for at least some progress
	if checkRunState != github.CheckRunFailure && checkRunState != github.CheckRunSuccess {
		ctx = workflow.WithRetryPolicy(ctx, temporal.RetryPolicy{
			MaximumAttempts: 3,
		})
	}

	// TODO: should we block here? maybe we can just make this async
	var resp activities.UpdateCheckRunResponse
	err := workflow.ExecuteActivity(ctx, n.Activity.UpdateCheckRun, request).Get(ctx, &resp)

	if err != nil {
		logger.Error(ctx, "updating check run", "err", err)
	}
}

func determineCheckRunState(workflowState *state.Workflow) github.CheckRunState {
	if workflowState.Apply == nil {
		switch workflowState.Plan.Status {
		case state.InProgressJobStatus, state.SuccessJobStatus:
			return github.CheckRunPending
		case state.FailedJobStatus:
			return github.CheckRunFailure
		}
	}

	switch workflowState.Apply.Status {
	case state.InProgressJobStatus:
		return github.CheckRunPending
	case state.SuccessJobStatus:
		return github.CheckRunSuccess
	}

	// this is a failure or rejection at this point
	return github.CheckRunFailure
}