package routes

import (
	"go_redis/module/testing"

	"github.com/labstack/echo"
)

// Routes is used to declare a route to controller that will be used to handle
func Routes(r *echo.Echo) *echo.Echo {
	r.GET("/create-list", testing.Test)
	r.GET("/delete/:key", testing.Deletes)
	r.GET("/detail/:key", testing.Detail)
	return r
}
