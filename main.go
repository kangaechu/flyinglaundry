package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kangaechu/flyinglaundry/tenki"
	"net/http"
	"os"
	"time"
)

func HandleRequest(ctx context.Context) (*string, error) {
	// TENKI_URL : https://tenki.jp/forecast/1/2/1400/1100/1hour.html
	url := os.Getenv("TENKI_URL")
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	tenki, err := tenki.NewTenki(res.Body)
	if err != nil {
		return nil, err
	}

	currentHour := time.Now().Hour()
	duration := 24
	flyingLaundry := tenki.CheckFlyingLaundry(currentHour, duration)

	SLACK_WEBHOOK_URL := os.Getenv("SLACK_WEBHOOK_URL")
	if flyingLaundry && SLACK_WEBHOOK_URL != "" {
		responseMessage := "@channel 洗濯物ふっとび注意"
		if err := sendSlack(SLACK_WEBHOOK_URL, responseMessage); err != nil {
			return nil, err
		}
	}
	returnString := fmt.Sprintf("Flying Laundry: %t", flyingLaundry)
	return &returnString, nil
}

func main() {
	lambda.Start(HandleRequest)
}

func sendSlack(url, message string) error {
	payload := map[string]interface{}{
		"text":       message,
		"link_names": 1,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message to Slack, status code: %d", resp.StatusCode)
	}

	return nil
}
