package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func HandlePaypackCallback(ctx echo.Context) error {
	req := ctx.Request()
	fmt.Printf("Receuved callback rquest: %v", req.Body)
	return ctx.String(200, "OK")
}
