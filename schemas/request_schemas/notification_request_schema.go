package requestschemas

type ServiceRequestSchema struct {
	ServiceName string   `json:"service_name"`
	WebHookUrls []string `json:"webhook_urls"`
}

type NotificationRequestSchema struct {
	ServiceName string `json:"service_name"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SlackNotificationRequestSchema struct {
	Message string `json:"text"`
}
