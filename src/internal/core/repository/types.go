package repository

import (
	"context"

	"github.com/Tomelin/desafio-stn/src/pkg/utils"
)

var timeout *utils.Timeout = &utils.Timeout{Millisecond: 5}

func getTimeout(ctx context.Context) {

	exists := ctx.Value("timeout")
	if exists != nil {
		*timeout = ctx.Value("timeout").(utils.Timeout)
	}

}
