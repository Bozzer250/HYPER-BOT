package webHandlers

import (
	"hyperbot/utils"
	publicPages "hyperbot/web/pages/public"

	"github.com/labstack/echo/v4"
)

func RenderLoginPage(c echo.Context) error {
	return utils.RenderVIews(c, publicPages.LoginPage())
}

func RenderLandingPage(c echo.Context) error {
	return utils.RenderVIews(c, publicPages.LandingPage())
}
