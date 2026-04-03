package validators

import (
	"fmt"

	rs "slack_notification_processor/schemas/request_schemas"
)

func ValidateNotificationRequestSchema(nrsObject *rs.NotificationRequestSchema) error {
	if nrsObject.Name == "" {
		return fmt.Errorf("field 'name' is required")
	}
	if nrsObject.Type == "" {
		return fmt.Errorf("field 'type' is required")
	}
	if nrsObject.Description == "" {
		return fmt.Errorf("field 'description' is required")
	}
	if nrsObject.ServiceName == "" {
		return fmt.Errorf("field 'service_name' is required")
	}

	return nil
}

func ValidateServiceSchema(scObject *rs.ServiceRequestSchema) error {
	if scObject.ServiceName == "" {
		return fmt.Errorf("field 'service_name' is required")
	}
	if len(scObject.WebHookUrls) == 0 {
		return fmt.Errorf("field 'webhook_urls' is required")
	}

	return nil
}
