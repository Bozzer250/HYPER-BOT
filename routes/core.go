package routes

import (
	"fmt"
	"hyperbot/controllers"
	"hyperbot/utils"
	publicComponents "hyperbot/web/components/public"

	"github.com/labstack/echo/v4"
)

func PurchaseNewAccount(ctx echo.Context) error {
	phone := ctx.FormValue("phone_number")
	referralCode := ctx.FormValue("referral_code")
	err := controllers.PurchaseNewAccount(phone, referralCode)
	if err != nil {
		return ctx.String(200, fmt.Sprintf("<div>There was an error purchasing the bot.</br>%v</br>Contact support to get help</div>", err))
	}
	return utils.RenderVIews(ctx, publicComponents.GoToDashboardButton())
}
