package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TautulliApiClient struct {
	baseUrl  string
	apiToken string
	http     *http.Client
}

type APIResponse[T TautulliUser | TautulliWatchTime] struct {
	Response struct {
		Result  string `json:"result"`
		Message string `json:"message"`
		Data    []T    `json:"data"`
	} `json:"response"`
}

type TautulliUser struct {
	UserId       int    `json:"user_id"`
	Username     string `json:"username"`
	FriendlyName string `json:"friendly_name"`
}

type TautulliWatchTime struct {
	QueryDays  int `json:"query_days"`
	TotalPlays int `json:"total_plays"`
	TotalTime  int `json:"total_time"`
}

type queryArgument struct {
	Name  string
	Value string
}

func NewClient(baseUrl string, apiToken string) *TautulliApiClient {
	return &TautulliApiClient{
		baseUrl:  baseUrl,
		apiToken: apiToken,
		http:     &http.Client{},
	}
}

func (c *TautulliApiClient) GetUsers() ([]TautulliUser, error) {
	data, err := c.doRequest("get_users", nil)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var users APIResponse[TautulliUser]
	if err := json.NewDecoder(data).Decode(&users); err != nil {
		return nil, err
	}
	if users.Response.Result != "success" {
		errorMessage := fmt.Sprintf("Could not successfully fetch users, returned result was: %s %s", users.Response.Result, users.Response.Message)
		return nil, errors.New(errorMessage)
	}

	return users.Response.Data, nil
}

func (c *TautulliApiClient) GetStats(userId int, timeframe int) (TautulliWatchTime, error) {
	queryArgs := []queryArgument{
		{Name: "user_id", Value: fmt.Sprint(userId)},
		{Name: "query_days", Value: fmt.Sprint(timeframe)},
	}
	data, err := c.doRequest("get_user_watch_time_stats", queryArgs)
	if err != nil {
		return TautulliWatchTime{}, err
	}
	defer data.Close()

	var watchTime APIResponse[TautulliWatchTime]
	if err := json.NewDecoder(data).Decode(&watchTime); err != nil {
		return TautulliWatchTime{}, err
	}
	if watchTime.Response.Result != "success" {
		errorMessage := fmt.Sprintf("Could not successfully fetch watch time for user %d, returned result was: %s %s", userId, watchTime.Response.Result, watchTime.Response.Message)
		return TautulliWatchTime{}, errors.New(errorMessage)
	}

	return watchTime.Response.Data[0], nil
}

func (c *TautulliApiClient) doRequest(command string, queryArgs []queryArgument) (io.ReadCloser, error) {
	url := c.baseUrl + "/api/v2?apikey=" + c.apiToken + "&cmd=" + command + c.formatArguments(queryArgs)

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

func (c *TautulliApiClient) formatArguments(arguments []queryArgument) string {
	if len(arguments) == 0 {
		return ""
	}

	var parsedArgs = ""
	for _, arg := range arguments {
		parsedArgs += "&" + url.QueryEscape(arg.Name) + "=" + url.QueryEscape(arg.Value)
	}
	return parsedArgs
}
