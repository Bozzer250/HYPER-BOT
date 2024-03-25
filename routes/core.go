package routes

import (
	"fmt"
	"hyperbot/controllers"
	"hyperbot/integrations"
	"hyperbot/utils"
	publicComponents "hyperbot/web/components/public"

	"github.com/labstack/echo/v4"
)

func PurchaseNewAccount(ctx echo.Context) error {
	phone := ctx.FormValue("phone_number")
	referralCode := ctx.FormValue("referral_code")
	packageName := ctx.FormValue("package_name")
	err := controllers.PurchaseNewAccount(phone, referralCode, packageName)
	if err != nil {
		return ctx.String(200, fmt.Sprintf("<div>There was an error purchasing the bot.</br>%v</br>Contact support to get help</div>", err))
	}
	return utils.RenderVIews(ctx, publicComponents.GoToDashboardButton())
}

func Test(c echo.Context) error {
	token, err := integrations.PaypackCashIn(100, "0788711226")
	if err != nil {
		return c.String(200, fmt.Sprintf("<div>There was an error cashing in.</br>%v</br>Contact support to get help</div>", err))
	}
	return c.String(200, fmt.Sprintf("<div>Token: %v</div>", token))
}
