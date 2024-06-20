package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
)

type RestAPI struct {
	Route *gin.Engine
	Group *gin.RouterGroup
}

func NewRestAPI(http2 bool, mode string) *RestAPI {

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	if http2 {
		r.UseH2C = true
	}

	return &RestAPI{
		Route: r,
		Group: r.Group("/api"),
	}
}

func (r *RestAPI) Run(handler http.Handler) error {

	r.Group.GET("/teste", r.GetAll)
	srv := http.Server{
		Addr:    ":8080",
		Handler: r.Route.Handler(),
	}

	http2.ConfigureServer(&srv, &http2.Server{})
	return srv.ListenAndServe()
}

func (ns *RestAPI) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, "result")
}
