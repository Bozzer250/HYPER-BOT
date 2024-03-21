package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleUserAuth(ctx echo.Context) error {
	phone := ctx.FormValue("phone_number")
	fmt.Println(phone)
	return ctx.String(http.StatusCreated, "<div>dial *182*7*1# and follow instructions to complete the topup</div>")
}
