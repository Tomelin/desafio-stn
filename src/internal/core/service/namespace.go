package service

import (
	"context"
	"github.com/Tomelin/desafio-stn/src/internal/core/repository"
	corev1 "k8s.io/api/core/v1"
)

type NamespaceServiceInterface interface {
	Count(ctx context.Context) (int, error)
	GetAll(ctx context.Context) ([]corev1.Namespace, error)
	GetByName(ctx context.Context, name string) (*corev1.Namespace, error)
	Create(ctx context.Context, request *corev1.Namespace) (*corev1.Namespace, error)
}

type ServiceNamespace struct {
	Repository repository.IRepositoryNamespace
}

func NewServiceNamespace(repo repository.IRepositoryNamespace) NamespaceServiceInterface {

	return &ServiceNamespace{
		Repository: repo,
	}
}

func (n *ServiceNamespace) Count(ctx context.Context) (int, error) {

	return n.Repository.Count(ctx)
}

func (n *ServiceNamespace) GetAll(ctx context.Context) ([]corev1.Namespace, error) {

	return n.Repository.GetAll(ctx)
}

func (n *ServiceNamespace) Create(ctx context.Context, request *corev1.Namespace) (*corev1.Namespace, error) {

	return n.Repository.Create(ctx, request)
}

func (n *ServiceNamespace) GetByName(ctx context.Context, name string) (*corev1.Namespace, error) {

	return n.Repository.GetByName(ctx, name)
}
