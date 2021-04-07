package testing

import (
	"go_redis/constant"
	"net/http"

	"github.com/labstack/echo"
)

// Test is used to Test Request
func Test(ctx echo.Context) error {
	response := new(constant.Response)
	a := new(Author)
	ctx.Bind(&a)

	value, err := a.Testing()

	if err != nil {
		response.Status = "405"
		response.Errors = err.Error()
		return ctx.JSON(http.StatusNotFound, response)
	}

	response.Status = "000"
	response.Message = "Success"
	response.Data = value

	return ctx.JSON(http.StatusOK, response)
}

// Test is used to Test Request
func Deletes(ctx echo.Context) error {
	response := new(constant.Response)

	key := ctx.Param("key")

	err := Delete(key)

	if err != nil {
		response.Status = "405"
		response.Errors = err.Error()
		return ctx.JSON(http.StatusNotFound, response)
	}

	response.Status = "000"
	response.Message = "Success"

	return ctx.JSON(http.StatusOK, response)
}

// Test is used to Test Request
func Detail(ctx echo.Context) error {
	response := new(constant.Response)
	key := ctx.Param("key")

	value, err := Find(key)

	if err != nil {
		response.Status = "405"
		response.Errors = err.Error()
		return ctx.JSON(http.StatusNotFound, response)
	}

	response.Status = "000"
	response.Message = "Success"
	response.Data = value

	return ctx.JSON(http.StatusOK, response)
}
