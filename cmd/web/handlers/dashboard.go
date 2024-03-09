package webHandlers

import (
	dashboardPages "hyperbot/cmd/web/pages/dashboard"
	"hyperbot/utils"

	"github.com/labstack/echo/v4"
)

func RenderDashboardLandingPage(c echo.Context) error {
	return utils.RenderVIews(c, dashboardPages.DashboardLandingPage())
}
