package demo

import (
	"demo-service/repository"
	"github.com/labstack/echo/v4"
)

type Route struct {
	route *repository.Repository
}

func Init(group *echo.Group, controller *repository.Repository) {
	r := &Route{controller}

	demo := group.Group("/demo")
	demo.GET("", r.Test)
}
