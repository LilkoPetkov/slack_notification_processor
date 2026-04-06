package handlers

import (
	"fmt"

	r "slack_notification_processor/requests"
	rs "slack_notification_processor/schemas/request_schemas"
	validators "slack_notification_processor/validators"

	l "slack_notification_processor/logger"
	vars "slack_notification_processor/vars"

	"github.com/gin-gonic/gin"
)

func MessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		notificationRequest := rs.NotificationRequestSchema{}

		if err := c.ShouldBindJSON(&notificationRequest); err != nil {
			l.Logger.Error(fmt.Sprintf("Invalid json error: %v", err))
			c.JSON(400, gin.H{"error": "invalid JSON"})
			return
		}

		err := validators.ValidateNotificationRequestSchema(&notificationRequest)
		if err != nil {
			l.Logger.Error(fmt.Sprintf("Validation error: %v", err.Error()))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if foundServiceWebhooks, ok := vars.SERVICES[notificationRequest.ServiceName]; ok {
			r.ProcessRequests(&notificationRequest, c, foundServiceWebhooks)
		} else {

			l.Logger.Error(fmt.Sprintf("Service does not exist: %v", err))
			c.JSON(400, gin.H{"error": fmt.Sprintf("service name '%s' does not exist", notificationRequest.ServiceName)})
		}
	}
}
