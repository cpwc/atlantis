package terraform

import (
	"github.com/runatlantis/atlantis/server/neptune/workflows/internal/github"
	"github.com/runatlantis/atlantis/server/neptune/workflows/internal/root"
)

type Request struct {
	Root         root.Root
	Repo         github.Repo
	DeploymentId string
	Revision     string
}