package infrastructure

import (
	"github.com/lucastomic/infrastructure-/internal/domain"
	"github.com/lucastomic/infrastructure-/internal/logging"
)

type InfrastructureService interface {
	// LocalDeploy takes the files of a generated WEB and deploys them in the local environment
	LocalDeploy(data domain.WebData) (domain.LocalDeployment, error)
}

type infrastructureService struct {
	logger logging.Logger
}

func New(l logging.Logger) InfrastructureService {
	return infrastructureService{l}
}

func (srv infrastructureService) LocalDeploy(data domain.WebData) (domain.LocalDeployment, error) {
	return domain.LocalDeployment{":3002"}, nil
}
