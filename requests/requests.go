package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	l "slack_notification_processor/logger"

	rs "slack_notification_processor/schemas/request_schemas"
)

type FailedRequests struct {
	URLs  []string
	Total int
}

type ProcessedRequests struct {
	URLs  []string
	Total int
}

func ProcessRequests(notificationRequest *rs.NotificationRequestSchema, c *gin.Context, foundServiceWebhooks []string) {
	nType, nName, nDesc := notificationRequest.Type, notificationRequest.Name, notificationRequest.Description

	if strings.ToLower(notificationRequest.Type) != "warning" {
		l.Logger.Info(fmt.Sprintf("Notification received. Type: %s - Name: %s - Description: %s", nType, nName, nDesc))
		c.JSON(http.StatusOK, gin.H{
			"msg": "notification logged, but will not be sent to channel as it is not warning",
		})
		return
	}

	failed := FailedRequests{}
	successful := ProcessedRequests{}

	message := fmt.Sprintf("Warning Notification Received\nName: %s\nDescription: %s\n", nName, nDesc)
	l.Logger.Info(message)
	payload, err := json.Marshal(rs.SlackNotificationRequestSchema{
		Message: message,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Cannot marshal schema: %v", notificationRequest),
		})
	}

	headers := map[string]string{"Content-Type": "application/json"}

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, url := range foundServiceWebhooks {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			_, err := httpRequest("post", u, headers, payload)
			if err != nil {
				l.Logger.Error(fmt.Sprintf("Error making HTTP request to webhook: %s: %v", url, err))
			}

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				l.Logger.Error(fmt.Sprintf("Error occurred while sending report: %v", err))
				failed.URLs = append(failed.URLs, u)
				failed.Total++
			} else {
				successful.URLs = append(successful.URLs, u)
				successful.Total++
			}
		}(string(url))
	}

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"successfully_processed": successful.URLs,
		"success_total":          successful.Total,
		"failed":                 failed.URLs,
		"failed_total":           failed.Total,
	})
}

func httpRequest(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(strings.ToUpper(method), url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: 25 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBody, fmt.Errorf("unexpected status code: %d - body: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
