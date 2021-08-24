package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/fatih/color"
	uuid "github.com/satori/go.uuid"
)

func (t *Task) InitClient(proxy *url.URL) { // Initializes HTTP Client and Cookiejar for each Task.
	defer handleError() // defers the handleError() function in case of a panic

	jar, err := cookiejar.New(nil) // initializing new cookieJar for task
	if err != nil {
		color.Red("[ERROR] InitClient.cookiejar FATAL ERROR CREATING COOKIE JAR: %s", err) // err handle
	}

	client := &http.Client{ // setting client fields - proxy, cookiejar
		Jar:       jar,
		Transport: &http.Transport{Proxy: http.ProxyURL(proxy), DisableKeepAlives: true}}

	t.Client = client // setting task client equal to generated client
}

func (t *Task) grabOrder() (*Orderinfo, error) {
	defer handleError()
	var body Orderinfo
	url := fmt.Sprintf("https://api.nike.com/orders/summary/v1/%s?locale=en_ca&country=CA&language=en-GB&email=%s", t.Orderid, t.Email)
	visitorId, _ := uuid.NewV4().MarshalText()

	headers := map[string][]string{
		"User-Agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"},
		"Authority":          {"api.nike.com"},
		"X-Nike-Visitid":     {"1"},
		"X-Nike-Visitorid":   {string(visitorId)},
		"Accept":             {"application/json"},
		"Nike-Api-Caller-Id": {"com.nike:sse.orders"},
		"Origin":             {"https://www.nike.com"},
		"Sec-Fetch-Site":     {"same-site"},
		"Sec-Fetch-Mode":     {"cors"},
		"Sec-Fetch-Dest":     {"empty"},
		"Referer":            {"https://www.nike.com/"},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		color.Red("[ERROR] grabOrder.NewRequest Error Initializing Request: %s\n", err)
		return nil, err
	}
	req.Header = headers

	res, err := t.Client.Do(req)
	if err != nil {
		color.Red("[ERROR] grabOrder.Client.Do Error Sending Request: %s\n", err)
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		jsonBody, err := io.ReadAll(res.Body)
		if err != nil {
			color.Red("[ERROR] grabOrder.ReadAll Error Reading Response Body: %s\n", err)
			return nil, err
		}

		color.HiCyan("[200] [%s] Successfully Grabbed Order Info.", t.Orderid)
		json.Unmarshal([]byte(jsonBody), &body)
	case 403:
		color.Red("[403] FORBIDDEN\n")
	case 404:
		color.Red("[404] Order Not Found.\n")
	case 500, 502, 503, 504:
		color.Blue("[%d] Server Error\n", res.StatusCode)
	default:
		color.Red("[%d] Unknown Response Received. Refer to Status Code for more information.\n", res.StatusCode)
	}
	return &body, nil
}

func handleError() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
