package main

import (
	"github.com/budiariyanto/cqrs-dummy/handler"
	"github.com/budiariyanto/cqrs-dummy/repo"
	"github.com/budiariyanto/cqrs-dummy/service"
	"github.com/budiariyanto/cqrs-dummy/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbDenormalize := common.NewDB("denormalizer")
	denormalizeRepo := repo.NewDenormalizeRepo(dbDenormalize)
	denormalizeService := service.NewDenormalizeService(denormalizeRepo)

	h := &handler.Handler{
		DenormalizeSvc: denormalizeService,
	}

	// Routes
	e.POST("/donation-create", h.CreateDonation)
	e.GET("/campaign-detail/:id", h.GetCampaignDetail)
	e.PUT("/campaign-detail/update-campaign-status", h.UpdateCampaignStatus)
	e.PUT("/campaign-detail/update-total-donation", h.UpdateTotalDonation)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
