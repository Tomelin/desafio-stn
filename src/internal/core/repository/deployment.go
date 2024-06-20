package repository

import (
	"context"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentRepositoryInterface interface {
	Count(ctx context.Context) (int, error)
	GetAll(ctx context.Context) ([]appsv1.Deployment, error)
	GetByName(ctx context.Context, name string) (*appsv1.Deployment, error)
	Create(ctx context.Context, request *appsv1.Deployment) (*appsv1.Deployment, error)
}

type DeploymentRepository struct {
	Client *kubernetes.Clientset
}

func NewDeploymentRepository(kubeClient *kubernetes.Clientset) DeploymentRepositoryInterface {

	return &DeploymentRepository{
		Client: kubeClient,
	}
}

func (n *DeploymentRepository) Count(ctx context.Context) (int, error) {

	getTimeout(ctx)

	ctx, cancel := context.WithTimeout(ctx, timeout.Millisecond*time.Millisecond)
	defer cancel()

	ns, err := n.Client.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})

	if err != nil {
		return 0, err
	}

	return len(ns.Items), nil
}

func (n *DeploymentRepository) GetAll(ctx context.Context) ([]appsv1.Deployment, error) {
	getTimeout(ctx)

	ctx, cancel := context.WithTimeout(ctx, timeout.Millisecond*time.Millisecond)
	defer cancel()

	ns, err := n.Client.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return ns.Items, nil
}

func (n *DeploymentRepository) GetByName(ctx context.Context, name string) (*appsv1.Deployment, error) {
	getTimeout(ctx)

	ctx, cancel := context.WithTimeout(ctx, timeout.Millisecond*time.Millisecond)
	defer cancel()

	ns, err := n.Client.AppsV1().Deployments("").Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	return ns, nil
}

func (n *DeploymentRepository) Create(ctx context.Context, request *appsv1.Deployment) (*appsv1.Deployment, error) {
	getTimeout(ctx)

	ctx, cancel := context.WithTimeout(ctx, timeout.Millisecond*time.Millisecond)
	defer cancel()

	ns, err := n.Client.AppsV1().Deployments("").Create(ctx, request, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return ns, nil
}
