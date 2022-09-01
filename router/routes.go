package router

import (
	otten "otten/internal/otten"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) (*echo.Echo, error) {

	e.GET("/otten", otten.GetTransaction)

	return e, nil
}
