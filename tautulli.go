package main

import (
	"io"
	"net/http"
)

type TautulliApiClient struct {
	baseUrl  string
	apiToken string
	http     *http.Client
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
	body, err := c.doRequest("get_users", nil)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	bodyBytes, err := io.ReadAll(body)
	println(string(bodyBytes))

	/*
		var users []User
		if err := json.NewDecoder(body).Decode(&users); err != nil {
			log.Printf("Could not unmarshal Tautulli users: %v", err)
			return nil, err
		}*/

	return nil, nil
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
