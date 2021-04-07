package main

import (
	"go_redis/routes"

	"github.com/labstack/echo"
)

func main() {
	e := routes.Routes(echo.New())

	e.Start(":6677")
}
