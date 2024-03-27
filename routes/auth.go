package routes

import (
	"fmt"
	"hyperbot/configs"
	"hyperbot/controllers"
	"hyperbot/utils"
	publicComponents "hyperbot/web/components/public"

	"github.com/labstack/echo/v4"
)

func HandleUserAuth(ctx echo.Context) error {
	phone := ctx.FormValue("phone_number")
	err := controllers.HandleUserAuth(phone)
	if err != nil {
		return ctx.String(200, fmt.Sprintf("<div>There was an error authenticating your request.</br>%v</br>Contact support to get help</div>", err))
	}
	// return ctx.String(http.StatusCreated, utils.RenderVIews(ctx, publicComponents.OtpForm(phone)))
	return utils.RenderVIews(ctx, publicComponents.OtpForm(phone, ""))
}

func VerifyOtp(ctx echo.Context) error {
	phone := ctx.FormValue("phone_number")
	code := ctx.FormValue("code")
	uid, err := controllers.VerifyAndUpdateOtp(phone, code)
	if err != nil {
		fmt.Printf("Error verifying OTP: %v\n", err)
		return utils.RenderVIews(ctx, publicComponents.OtpForm(phone, "There was an error verifying your OTP. Please try again"))
	}
	Store := configs.SessionStore
	session, _ := Store.Get(ctx.Request(), "hyper-bots")
	// Set some session values.
	session.Values["uid"] = uid
	errr := Store.Save(ctx.Request(), ctx.Response(), session)
	if errr != nil {
		fmt.Printf("Error saving session: %v\n", errr)
	}
	return utils.RenderVIews(ctx, publicComponents.GoToDashboardButton())
}
