package webserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/net/http2"
)

type RestAPI struct {
	Route *gin.Engine
	Group *gin.RouterGroup
}

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})

var routerPath *gin.RouterGroup

func init() {
	prometheus.MustRegister(cpuTemp)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func NewRestAPI(http2 bool, mode string) *RestAPI {

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	if http2 {
		r.UseH2C = true
	}

	routerGroupPath := fmt.Sprintf("/api")
	routerPath = r.Group(routerGroupPath)

	r.GET("/metrics", prometheusHandler())

	// Set swagger
	routerPath.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	routerPath.GET("/docs/swagger", func(c *gin.Context) {
		c.Redirect(301, fmt.Sprintf("%s/docs/swagger/index.html", routerGroupPath))
	})
	routerPath.GET("/docs", func(c *gin.Context) {

		c.Redirect(301, fmt.Sprintf("%s/docs/swagger/index.html", routerGroupPath))
	})
	routerPath.GET("/", func(c *gin.Context) {
		c.Redirect(301, fmt.Sprintf("%s/docs/swagger/index.html", routerGroupPath))
	})

	rest := &RestAPI{
		Route: r,
	}

	routerPath.Use(rest.MiddlewareAuthorization)
	rest.Group = routerPath

	return rest
}

func (r *RestAPI) Run(handler http.Handler, listen, port string) error {

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", listen, port),
		Handler: r.Route.Handler(),
	}

	http2.ConfigureServer(&srv, &http2.Server{})
	return srv.ListenAndServe()
}

func (s *RestAPI) MiddlewareAuthorization(c *gin.Context) {

	if c.GetHeader("Authorization") != "ZGVzYWZpby1zdG4K" {
		fmt.Println("Authorization is not valid")
	}

	c.Next()
}
