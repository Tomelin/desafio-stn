package webserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

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

func (ns *NamespaceWebServer) handlers(router *gin.RouterGroup) {
	router.GET("/namespace", ns.GetAll)
	router.GET("/namespace/search", ns.GetByName)
	router.GET("/namespace/count", ns.Count)
}

// NamespaceCount 	godoc
// @Summary 				count namespaces
// @Description 		count the number of namespaces
// @Tags namespace
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /namespace [get]
func (ns *NamespaceWebServer) Count(c *gin.Context) {

	result, err := ns.Service.Count(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, result)

}
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

func (ns *NamespaceWebServer) GetByName(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	fmt.Println()

	if len(c.Request.URL.Query()) > 0 {
		query := c.Request.URL.Query()
		if name, exists := query["name"]; exists {
			fmt.Println(exists, name)
			result, err := ns.Service.GetByName(ctx, name[0])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				c.Abort()
			}

			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "key name not found in url",
			})
			c.Abort()
		}
	}

}
