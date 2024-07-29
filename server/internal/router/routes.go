package router

import (
	"fmt"
	"net/http"
	"stockanalyzer/internal/container"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	mainGroup *echo.Group
}

func NewRouter(e *echo.Echo, container *container.Container) *Router {
	router := &Router{}
	router.initializeRouter(e, container)
	return router
}

func (r *Router) initializeRouter(e *echo.Echo, container *container.Container) {
	setCORSConfig(e, container)
	r.mainGroup = e.Group("/api/v1")
}

func setCORSConfig(e *echo.Echo, container *container.Container) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
		AllowOrigins:                             []string{"*"},
		AllowHeaders: []string{
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderAcceptEncoding,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		MaxAge: 86400,
	}))
}

func (r *Router) RegisterRoutes(domain string, routes []Route) {
	fmt.Println("Registering routes for domain", domain)
	//Set the group for the domain
	group := r.mainGroup.Group(domain)

	//Register the routes
	for _, route := range routes {

		group.Add(route.Method, route.Path, route.Handler)
	}
}

func (r *Router) SetRoutes(path string, routes []Route) error {
	r.RegisterRoutes(path, routes)
	return nil
}

func (r *Router) BuildRoute(method string, path string, handler echo.HandlerFunc) Route {
	return Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
