package routes

import (
	"github.com/gin-gonic/gin"
	"slack_notification_processor/handlers"
)

func SetupRoutes(r *gin.Engine, v1Group *gin.RouterGroup) {
	v1Group.POST("/register-service", handlers.ServiceHandler())
	v1Group.POST("/send-notification", handlers.MessageHandler())
}
