# Slack Notification Processor

A Go-based service that receives notifications via HTTP and forwards them to Slack webhooks.

## Slack setup
1. Manual setup of channels
  - Create new channel in Slack
  - Click on Admin -> Apps & Workflows
  - Search for `Incoming Webhook`
  - Add the webhook to the channel
  - Use the webhook URL when registering a service.
2. Precreated channel links to join for testing:
  - https://join.slack.com/share/enQtMTA4NDA3NDUwMDY4MzQtNGM5NTZlZWQ5Njc2YTNkMmZhZTkwYzU3Y2ZmMjE5MDlmYjJmZDBlMzc5N2YxYWIzMzNhYzNjZWVkYWU1N2Q0Mg
  - https://join.slack.com/share/enQtMTA4NDA3NjIwMTc0NDItYjczNjBmYTQzMGI1OWRjYTg3MzlmNzEzM2Q2NmQ2MjI2MTAwYTc5Zjg3YWYzNWNhMTczZjIwYzJjYWU4NjYyYQ


## Running with Docker

1. Build the Docker image:
```bash
docker build -t slack-notification-processor .
```

2. Run the Docker container:
   
```bash
docker run -d -p 8081:8081 --name slack-notifier slack-notification-processor
```

The service will be available at http://localhost:8081.

## API Usage

The service requires you to first register a service name (simulating at least basic type of authentication) with one or more Slack webhook URLs. Once registered, you can send notifications using that service name.

### 1. Register a Service

Use the `/v1/register-service` endpoint to map a `service_name` to multiple Slack webhook URLs.

**Sample CURL Request:**
```bash
curl -X POST http://localhost:8081/v1/register-service \
     -H "Content-Type: application/json" \
     -d '{
       "service_name": "test_service",
       "webhook_urls": [
         "https://hooks.slack.com/services/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",
         "https://hooks.slack.com/services/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"
       ]
     }'
```

### 2. Send a Notification

Use the `/v1/send-notification` endpoint to send a message to all webhooks associated with a registered `service_name`.

**Sample CURL Request:**
```bash
curl -X POST http://localhost:8081/v1/send-notification \
     -H "Content-Type: application/json" \
     -d '{
       "service_name": "test_service",
       "type": "warning",
       "name": "Backup Failer",
       "description": "The Backup failed due to database problem"
     }'
```
