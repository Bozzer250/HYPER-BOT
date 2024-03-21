package webHandlers

import (
	"hyperbot/utils"
	dashboardPages "hyperbot/web/pages/dashboard"

	"github.com/labstack/echo/v4"
)

func RenderDashboardLandingPage(c echo.Context) error {
	return utils.RenderVIews(c, dashboardPages.DashboardLandingPage())
}
