package cmd

import (
	"ewallet-framework/helpers"
	"ewallet-framework/internal/api"
	"ewallet-framework/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}
	r := gin.Default()

	r.GET("/health", healthcheckAPI.HealthcheckHandlerHTTP)
	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
