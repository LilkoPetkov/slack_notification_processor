package validators

import (
	"fmt"

	rs "slack_notification_processor/schemas/request_schemas"
	vars "slack_notification_processor/vars"
)

func ValidateServiceSchema(scObject *rs.ServiceRequestSchema) error {
	if scObject.ServiceName == "" {
		return fmt.Errorf("field 'service_name' is required")
	}
	if len(scObject.WebHookUrls) == 0 {
		return fmt.Errorf("field 'webhook_urls' is required")
	}
	if _, ok := vars.SERVICES[scObject.ServiceName]; ok {
		return fmt.Errorf("service '%s' already exists", scObject.ServiceName)
	}

	return nil
}
