package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kangaechu/flyinglaundry/tenki"
	"io"
	"net/http"
	"os"
	"time"
)

func HandleRequest(_ context.Context) (*string, error) {
	// TENKI_URL : https://tenki.jp/forecast/1/2/1400/1100/1hour.html
	url := os.Getenv("TENKI_URL")
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	tnk, err := tenki.NewTenki(res.Body)
	if err != nil {
		return nil, err
	}

	currentHour := time.Now().Hour()
	duration := 24
	flyingLaundry := tnk.CheckFlyingLaundry(currentHour, duration)

	SlackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	if flyingLaundry && SlackWebhookUrl != "" {
		responseMessage := "@channel 🪽洗濯物ふっとび注意🌪️"
		if err := sendSlack(SlackWebhookUrl, responseMessage); err != nil {
			return nil, err
		}
	}
	returnString := fmt.Sprintf("Flying Laundry: %t", flyingLaundry)
	println(returnString)
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
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message to Slack, status code: %d", resp.StatusCode)
	}

	return nil
}
