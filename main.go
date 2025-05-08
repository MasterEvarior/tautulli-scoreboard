package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := getEnvVarWithDefault("GIZE_PORT", ":8080")

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalf("Could not start the server because of the following issue: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	API_TOKEN := getEnvVar("TS_API_TOKEN")
	BASE_URL := getEnvVar("TS_BASE_URL")

	tautulliClient := NewClient(BASE_URL, API_TOKEN)
	users, err := tautulliClient.GetUsers()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	log.Printf("Fetched %d users", len(users))

	for _, user := range users {
		watchTime, _ := tautulliClient.GetStats(user.UserId, 7)
		println(user.FriendlyName)
		println(watchTime.QueryDays)
		println(watchTime.TotalPlays)
		println(watchTime.TotalTime)
		println("----------------------")
	}

}

func getEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Environment variable '%s' was not defined", name)
	}
	return value
}

func getEnvVarWithDefault(name string, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	return value
}
