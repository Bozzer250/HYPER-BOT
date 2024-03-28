package webHandlers

import (
	"hyperbot/configs"
	"hyperbot/models"
	"hyperbot/utils"
	dashboardPages "hyperbot/web/pages/dashboard"

	"github.com/labstack/echo/v4"
)

func RenderDashboardLandingPage(ctx echo.Context) error {
	Store := configs.SessionStore
	session, _ := Store.Get(ctx.Request(), "hyper-bots")
	userId := session.Values["uid"].(string)
	totalAssets, _ := models.GetSumOfAssetsByUserId(userId)
	return utils.RenderVIews(ctx, dashboardPages.DashboardLandingPage(totalAssets))
}

func RenderProfilePage(ctx echo.Context) error {
	Store := configs.SessionStore
	session, _ := Store.Get(ctx.Request(), "hyper-bots")
	userId := session.Values["uid"].(string)
	user := models.GetUserById(userId)
	return utils.RenderVIews(ctx, dashboardPages.ProfilePage(user))
}
