package main

import (
	"encoding/json"
	"net/http"
)

type TautulliApiClient struct {
	baseUrl  string
	apiToken string
	http     *http.Client
}

type APIResponse struct {
	Response struct {
		Result string          `json:"result"`
		Data   json.RawMessage `json:"data"`
	} `json:"response"`
}

type User struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
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

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func getStats() {

}

func (c *TautulliApiClient) doRequest(command string, arguments []string) (json.RawMessage, error) {
	url := c.baseUrl + "/api/v2?apikey=" + c.apiToken + "&cmd=" + command

	println(url)

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

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Response.Data, nil
}

func (c *TautulliApiClient) formatArguments(arguments []string) string {
	if arguments == nil || len(arguments) == 0 {
		return ""
	}

	return ""
}
