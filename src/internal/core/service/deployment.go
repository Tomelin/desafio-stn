package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Tomelin/desafio-stn/src/internal/core/repository"
	appsv1 "k8s.io/api/apps/v1"
)

type DeploymentServiceInterface interface {
	Count(ctx context.Context) (int, error)
	GetAll(ctx context.Context) ([]appsv1.Deployment, error)
	GetByName(ctx context.Context, name string) (*appsv1.Deployment, error)
	Create(ctx context.Context, request *appsv1.Deployment) (*appsv1.Deployment, error)
}

type DeploymentService struct {
	Repository repository.DeploymentRepositoryInterface
}

func NewDeploymentService(repo repository.DeploymentRepositoryInterface) DeploymentServiceInterface {

	return &DeploymentService{
		Repository: repo,
	}
}

func (n *DeploymentService) Count(ctx context.Context) (int, error) {

	return n.Repository.Count(ctx)
}

func (n *DeploymentService) GetAll(ctx context.Context) ([]appsv1.Deployment, error) {

	return n.Repository.GetAll(ctx)
}

func (n *DeploymentService) Create(ctx context.Context, request *appsv1.Deployment) (*appsv1.Deployment, error) {

	if request == nil {
		return nil, errors.New("the object deployment cannot be empty or nil")
	}

	if request.Namespace == "" {
		return nil, errors.New("the deployment metadata.namespace cannot be empty")
	}

	if request.Name == "" {
		return nil, errors.New("the deployment metadata.name cannot be empty")
	}

	if request.Kind != "Deployment" {
		return nil, errors.New("the kind should be Deployment")
	}

	if request.Spec.Template.Spec.Containers == nil {
		return nil, errors.New("the spec.containers should not be empty")
	}

	_, err := n.GetByName(ctx, request.Name)

	if err != nil {
		if err.Error() != fmt.Sprintf("deployments \"%s\" not found", request.Name) {
			return nil, err
		}
	}

	return n.Repository.Create(ctx, request)
}

func (n *DeploymentService) GetByName(ctx context.Context, name string) (*appsv1.Deployment, error) {

	return n.Repository.GetByName(ctx, name)
}
