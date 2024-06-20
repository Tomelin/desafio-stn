package webserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	corev1 "k8s.io/api/core/v1"

	"github.com/Tomelin/desafio-stn/src/internal/core/service"
)

type NamespaceWebServer struct {
	Service service.NamespaceServiceInterface
}

func NewNamespaceWebServer(routerGroup *gin.RouterGroup, svc service.NamespaceServiceInterface) HandlerWebServer {

	ns := &NamespaceWebServer{
		Service: svc,
	}
	ns.handlers(routerGroup)
	return ns
}
func (ns *NamespaceWebServer) handlers(routerGroup *gin.RouterGroup) {

	routerGroup.GET("/namespace", ns.GetAll)
	routerGroup.GET("/namespace/:name", ns.GetByName)
	routerGroup.GET("/namespace/count", ns.Count)
	routerGroup.POST("/namespace", ns.Create)
}

// NamespaceCount 	godoc
// @Summary 				count namespaces
// @Description 		count the number of namespaces
// @Tags namespace
// @Accept json
// @Produce json
// @Success 200 {number} int
// @Router /api/namespace/count [get]
func (ns *NamespaceWebServer) Count(c *gin.Context) {

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

// NamespaceCount 	godoc
// @Summary 				list all namespaces
// @Description 		List all namespaces from Kubernetes
// @Tags namespace
// @Accept json
// @Produce json
// @Success 200 {object} []corev1.Namespace
// @Router /api/namespace [get]
func (ns *NamespaceWebServer) GetAll(c *gin.Context) {
	result, err := ns.Service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, result)
}

// NamespaceGetByName 	godoc
// @Summary 				get namespace by name
// @Description 		get namespace by name from Kubernetes
// @Tags namespace
// @Param name path string true "found"
// @Accept json
// @Produce json
// @Success 200 {object} corev1.Namespace
// @Router /api/namespace/{name} [get]
func (ns *NamespaceWebServer) GetByName(c *gin.Context) {
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

// NamespaceCreate 	godoc
// @Summary 				create a new namespace
// @Description 		create a new namespace in Kubernetes
// @Tags namespace
// @Param       request body corev1.Namespace true "Create a namespace"
// @Accept json
// @Produce json
// @Success 200 {object} corev1.Namespace
// @Failure     403 {object} string
// @Failure     500 {object} string
// @Router /api/namespace [post]
func (ns *NamespaceWebServer) Create(c *gin.Context) {

	var object *corev1.Namespace
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
