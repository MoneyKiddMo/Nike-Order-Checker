package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type Webhook struct {
	Content   interface{} `json:"content"`
	Embeds    []Embeds    `json:"embeds"`
	Username  string      `json:"username"`
	AvatarURL string      `json:"avatar_url"`
}
type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type Thumbnail struct {
	URL string `json:"url"`
}
type Embeds struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Color       int       `json:"color"`
	Fields      []Fields  `json:"fields"`
	Footer      Footer    `json:"footer"`
	Timestamp   time.Time `json:"timestamp"`
	Thumbnail   Thumbnail `json:"thumbnail"`
}

func sendWebhook(hookUrl string, errInt int, successInt int) {
	defer handleError()

	successCount := fmt.Sprintf("%d", successInt)
	errorCount := fmt.Sprintf("%d", errInt)

	Webhook := Webhook{
		Embeds: []Embeds{
			Embeds{
				Title: "Nike Order Recap",
				Color: 0,
				Fields: []Fields{
					Fields{
						Name:  "Success Count",
						Value: successCount,
					},
					Fields{
						Name:  "Error Count",
						Value: errorCount,
					},
				},
				Timestamp: time.Now(),
				Footer: Footer{
					Text:    "Nike Order Checker | jc2#9899",
					IconURL: "https://pbs.twimg.com/profile_images/1407939286268317696/92L0QtAa_400x400.jpg",
				},
			},
		},
		Username:  "Nike Order Checker",
		AvatarURL: "https://cdn.freebiesupply.com/logos/large/2x/nike-4-logo-png-transparent.png",
	}

	webhookData, err := json.Marshal(Webhook)
	if err != nil {
		color.Red("[ERR] webhook.Marshal ERROR MARSHALLING WEBHOOK JSON DATA: %s\n", err)
	}

	req, err := http.NewRequest("POST", hookUrl, bytes.NewBuffer(webhookData))
	if err != nil {
		color.Red("[ERR] webhook.Post.NewRequest ERROR INIT POST REQUEST: %s\n", err)
	}
	req.Close = true

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		color.Red("[ERR] webhook.Client.Do ERROR SENDING POST REQUEST FOR WEBHOOK: %s\n", err)
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		color.Green("[200] Sent Recap Webhook Successfully!")
		time.Sleep(time.Second * 1)
	case 204:
		color.Green("[200] Sent Recap Webhook Successfully")
		time.Sleep(time.Second * 1)
	case 403:
		color.Red("[403] Error Sending Webhook: FORBIDDEN")
	case 401:
		color.Red("[401] Error Sending Webhook: Unauthorized")
	default:
		color.Yellow("[%d] Unknown Response Sending Webhook. Refer to Status code For More Info.", res.StatusCode)
	}
}
