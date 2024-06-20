package entity

// import (
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// )

import (
	"errors"
)

type Namespace struct {
	Name   string            "json:'name'"
	Labels map[string]string "json:'labels'"
}

func (n *Namespace) Validate() error {

	if n.Name == "" {
		return errors.New("namesapce cannot be empty")
	}

	return nil
}
