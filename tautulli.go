package main

type User struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}

type WatchTime struct {
	QueryDays  int `json:"query_days"`
	TotalPlays int `json:"total_plays"`
	TotalTime  int `json:"total_time"`
}

func getUsers() {

}

func getStats() {

}
