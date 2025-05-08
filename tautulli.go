package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TautulliApiClient struct {
	baseUrl  string
	apiToken string
	http     *http.Client
}

type APIResponse[T User | WatchTime] struct {
	Response struct {
		Result  string `json:"result"`
		Message string `json:"message"`
		Data    []T    `json:"data"`
	} `json:"response"`
}

type User struct {
	UserId       int    `json:"user_id"`
	Username     string `json:"username"`
	FriendlyName string `json:"friendly_name"`
}

type WatchTime struct {
	QueryDays  int `json:"query_days"`
	TotalPlays int `json:"total_plays"`
	TotalTime  int `json:"total_time"`
}

func NewClient(baseUrl string, apiToken string) *TautulliApiClient {
	return &TautulliApiClient{
		baseUrl:  baseUrl,
		apiToken: apiToken,
		http:     &http.Client{},
	}
}

func (c *TautulliApiClient) GetUsers() ([]User, error) {
	data, err := c.doRequest("get_users", nil)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var users APIResponse[User]
	if err := json.NewDecoder(data).Decode(&users); err != nil {
		return nil, err
	}
	if users.Response.Result != "success" {
		errorMessage := fmt.Sprintf("Could not successfully fetch users, returned result was: %s %s", users.Response.Result, users.Response.Message)
		return nil, errors.New(errorMessage)
	}

	return users.Response.Data, nil
}

func getStats() {

}

func (c *TautulliApiClient) doRequest(command string, arguments []string) (io.ReadCloser, error) {
	url := c.baseUrl + "/api/v2?apikey=" + c.apiToken + "&cmd=" + command

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (c *TautulliApiClient) formatArguments(arguments []string) string {
	if arguments == nil || len(arguments) == 0 {
		return ""
	}

	return ""
}
