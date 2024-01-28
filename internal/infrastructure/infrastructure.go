package infrastructure

import (
	"github.com/lucastomic/infrastructure-/internal/logging"
	"github.com/lucastomic/infrastructure-/internal/types"
)

type InfrastructureService interface {
	// LocalDeploy takes the files of a generated WEB and deploys them in the local environment
	LocalDeploy(data types.WebData) types.LocalDeployment
}

type infrastructureService struct {
	logger logging.Logger
}

func (srv infrastructureService) LocalDeploy(data types.WebData) types.LocalDeployment {
}
