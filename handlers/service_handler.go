package handlers

import (
	"fmt"
	"net/http"

	rs "slack_notification_processor/schemas/request_schemas"
	validators "slack_notification_processor/validators"

	l "slack_notification_processor/logger"
	vars "slack_notification_processor/vars"

	"github.com/gin-gonic/gin"
)

func ServiceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceRequest := rs.ServiceRequestSchema{}

		if err := c.ShouldBindJSON(&serviceRequest); err != nil {
			l.Logger.Error(fmt.Sprintf("Invalid json error: %v", err))
			c.JSON(400, gin.H{"error": "invalid JSON"})
			return
		}

		err := validators.ValidateServiceSchema(&serviceRequest)
		if err != nil {
			l.Logger.Error(fmt.Sprintf("Validation error: %v", err.Error()))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		vars.SERVICES[serviceRequest.ServiceName] = serviceRequest.WebHookUrls
		l.Logger.Error(fmt.Sprintf("New service created: '%s'", serviceRequest.ServiceName))

		c.JSON(http.StatusCreated, gin.H{
			"msg": fmt.Sprintf("service created: '%s'", serviceRequest.ServiceName),
		})
	}
}
