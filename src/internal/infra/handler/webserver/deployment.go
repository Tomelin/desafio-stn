package webserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	appsv1 "k8s.io/api/apps/v1"

	"github.com/Tomelin/desafio-stn/src/internal/core/service"
)

type DeploymentWebServer struct {
	Service service.DeploymentServiceInterface
}

func NewDeploymentWebServer(routerGroup *gin.RouterGroup, svc service.DeploymentServiceInterface) HandlerWebServer {

	ns := &DeploymentWebServer{
		Service: svc,
	}
	ns.handlers(routerGroup)
	return ns
}
func (ns *DeploymentWebServer) handlers(routerGroup *gin.RouterGroup) {

	routerGroup.GET("/deployment", ns.GetAll)
	routerGroup.GET("/deployment/:name", ns.GetByName)
	routerGroup.GET("/deployment/count", ns.Count)
	routerGroup.POST("/deployment", ns.Create)
}

// DeploymentCount 	godoc
// @Summary 				count deployments
// @Description 		count the number of deployments
// @Tags deployment
// @Accept json
// @Produce json
// @Success 200 {number} int
// @Router /api/deployment/count [get]
func (ns *DeploymentWebServer) Count(c *gin.Context) {

	result, err := ns.Service.Count(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}

	fmt.Println(result)
	c.JSON(http.StatusOK, result)

}

// DeploymentCount 	godoc
// @Summary 				list all deployments
// @Description 		List all deployments from Kubernetes
// @Tags deployment
// @Accept json
// @Produce json
// @Success 200 {object} []appsv1.Deployment
// @Router /api/deployment [get]
func (ns *DeploymentWebServer) GetAll(c *gin.Context) {
	result, err := ns.Service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, result)
}

// DeploymentGetByName 	godoc
// @Summary 				get deployment by name
// @Description 		get deployment by name from Kubernetes
// @Tags deployment
// @Param name path string true "found"
// @Accept json
// @Produce json
// @Success 200 {object} appsv1.Deployment
// @Router /api/deployment/{name} [get]
func (ns *DeploymentWebServer) GetByName(c *gin.Context) {
	name := c.Param("name")

	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "key name not found in url",
		})
		c.Abort()
	}

	result, err := ns.Service.GetByName(ctx, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, result)
}

// DeploymentCreate 	godoc
// @Summary 				create a new deployment
// @Description 		create a new deployment in Kubernetes
// @Tags deployment
// @Param       request body appsv1.Deployment true "Create a deployment"
// @Accept json
// @Produce json
// @Success 200 {object} appsv1.Deployment
// @Failure     403 {object} string
// @Failure     500 {object} string
// @Router /api/deployment [post]
func (ns *DeploymentWebServer) Create(c *gin.Context) {

	var object *appsv1.Deployment

	if err := c.ShouldBindWith(&object, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	object, err := ns.Service.Create(context.Background(), object)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": object})

}
