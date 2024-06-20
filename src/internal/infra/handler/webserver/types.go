package webserver

import (
	"context"

	"github.com/gin-gonic/gin"
)

type HandlerWebServer interface {
	Count(c *gin.Context)
	GetAll(c *gin.Context)
	GetByName(c *gin.Context)
}

var ctx context.Context = context.Background()
