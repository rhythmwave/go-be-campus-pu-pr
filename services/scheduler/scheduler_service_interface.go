package scheduler

import (
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
)

// SchedulerServiceInterface SchedulerServiceInterface
type SchedulerServiceInterface interface {
	AutoSetLeaveActive()
	AutoSetActiveSemester()
}

// NewSchedulerService function to bind myService to SchedulerServiceInterface
func NewSchedulerService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) SchedulerServiceInterface {
	return &schedulerService{
		repoCtx,
		infraCtx,
	}
}
