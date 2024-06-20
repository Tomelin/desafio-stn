package repository

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type IRepositoryNamespace interface {
	Count(ctx context.Context) (int, error)
	GetAll(ctx context.Context) ([]corev1.Namespace, error)
	GetByName(ctx context.Context, name string) (*corev1.Namespace, error)
	Create(ctx context.Context, request *corev1.Namespace) (*corev1.Namespace, error)
}

type RepositoryNamespace struct {
	Client *kubernetes.Clientset
}

func NewRepositoryNamespace(kubeClient *kubernetes.Clientset) IRepositoryNamespace {

	return &RepositoryNamespace{
		Client: kubeClient,
	}
}

func (n *RepositoryNamespace) Count(ctx context.Context) (int, error) {

	getTimeout(ctx)

	ctx, cancel := context.WithTimeout(ctx, timeout.Millisecond*time.Millisecond)
	defer cancel()

	ns, err := n.Client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})

	if err != nil {
		return 0, err
	}

	return len(ns.Items), nil
}

func (n *RepositoryNamespace) GetAll(ctx context.Context) ([]corev1.Namespace, error) {
	getTimeout(ctx)

	ctx, cancel := context.WithTimeout(ctx, timeout.Millisecond*time.Millisecond)
	defer cancel()

	ns, err := n.Client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return ns.Items, nil
}

func (n *RepositoryNamespace) GetByName(ctx context.Context, name string) (*corev1.Namespace, error) {
	getTimeout(ctx)

	ctx, cancel := context.WithTimeout(ctx, timeout.Millisecond*time.Millisecond)
	defer cancel()

	ns, err := n.Client.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	return ns, nil
}

func (n *RepositoryNamespace) Create(ctx context.Context, request *corev1.Namespace) (*corev1.Namespace, error) {
	getTimeout(ctx)

	ctx, cancel := context.WithTimeout(ctx, timeout.Millisecond*time.Millisecond)
	defer cancel()

	ns, err := n.Client.CoreV1().Namespaces().Create(ctx, request, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return ns, nil
}
