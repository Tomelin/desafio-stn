package kube

import (
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var config *rest.Config
var clientSet *kubernetes.Clientset

func NewKubeConnection(kubeConfigFilePath string) (*kubernetes.Clientset, error) {

	if len(kubeConfigFilePath) == 0 {
		c, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
		config = c
	} else {
		//load from a kube config
		var kubeconfig string

		if kubeConfigFilePath == "" {
			if home := homedir.HomeDir(); home != "" {
				kubeconfig = filepath.Join(home, ".kube", "config")
			}
		} else {
			kubeconfig = kubeConfigFilePath
		}

		c, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
		config = c
	}

	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return cs, nil

}
