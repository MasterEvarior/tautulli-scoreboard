package main

import (
	"log"
	"os"
)

func getEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Environment variable '%s' was not defined", name)
	}
	return value
}

func main() {
	API_TOKEN := getEnvVar("TS_API_TOKEN")
	BASE_URL := getEnvVar("TS_BASE_URL")

	tautulliClient := NewClient(BASE_URL, API_TOKEN)
	users, err := tautulliClient.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		println(u.UserId)
		println(u.Username)
		println(u.FriendlyName)
	}
}
