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
		log.Printf("Could not users because of an error: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched %d users", len(users))

	var scoreboard []WatchTime
	for _, user := range users {
		watchTime, err := tautulliClient.GetStats(user.UserId, 7)
		if err != nil {
			log.Printf("Could not fetch status for user %s because of an error: %v", user.FriendlyName, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if watchTime.TotalTime > 0 {
			scoreboard = append(scoreboard, watchTime)
			log.Printf("Added %s to the scoreboard", user.FriendlyName)
			continue
		}

		log.Printf("Not adding %s to the scoreboard, because the watchtime was 0 or less", user.FriendlyName)
	}

	log.Printf("Fetched %d watch times for the scoreboard", len(scoreboard))
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
