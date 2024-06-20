package entity

import (
	"errors"
)

type Namespace struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

func (n *Namespace) Validate() error {

	if n.Name == "" {
		return errors.New("namesapce cannot be empty")
	}

	return nil
}
