package main

import (
	"fmt"

	l "slack_notification_processor/logger"

	"github.com/gin-gonic/gin"
	"slack_notification_processor/routes"
)

func main() {
	defer l.Close()

	router := gin.Default()
	v1Group := router.Group("/v1")

	routes.SetupRoutes(router, v1Group)

	l.Logger.Info("Starting server on :8081")
	if err := router.Run(":8081"); err != nil {
		l.Logger.Error(fmt.Sprintf("could not run server: %v", err))
	}
}
