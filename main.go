package main

import (
	"log"
	"net/http"
	"slices"
	"sort"
	"text/template"

	_ "embed"
)

func main() {
	PORT := GetEnvVarWithDefault("TS_PORT", ":8080")

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalf("Could not start the server because of the following issue: %v", err)
	}
}

//go:embed index.html
var indexTemplate string

type templateData struct {
	Title  string
	Footer string
	Users  []User
}

type User struct {
	Name      string
	WatchTime float64
	Plays     int
}

func handler(w http.ResponseWriter, r *http.Request) {
	API_TOKEN := GetEnvVar("TS_API_TOKEN")
	BASE_URL := GetEnvVar("TS_BASE_URL")

	tautulliClient := NewClient(BASE_URL, API_TOKEN)
	users, err := tautulliClient.GetUsers()
	if err != nil {
		log.Printf("Could not users because of an error: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched %d users", len(users))

	var scoreboardUsers []User
	for _, user := range users {
		watchTime, err := tautulliClient.GetStats(user.UserId, getTimeframe(r))
		if err != nil {
			log.Printf("Could not fetch status for user %s because of an error: %v", user.FriendlyName, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if watchTime.TotalTime <= 0 {
			log.Printf("Not adding %s to the scoreboard, because the watchtime was 0 or less", user.FriendlyName)
			continue
		}

		scoreboardUsers = append(scoreboardUsers, User{
			Name:      user.FriendlyName[0:3],
			WatchTime: toHours(watchTime.TotalTime),
			Plays:     watchTime.TotalPlays,
		})
		log.Printf("Added %s to the scoreboard", user.FriendlyName)

	}

	log.Printf("Fetched %d watch times for the scoreboard", len(scoreboardUsers))

	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	err = tmpl.Execute(w, getTemplateData(scoreboardUsers))
	if err != nil {
		log.Printf("Could not render status because of an error: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func getTimeframe(r *http.Request) int {
	allowedTimeframes := []int{1, 7, 30, 365}

	if slices.Contains(allowedTimeframes, 7) {
		return 7
	}

	return 7
}

func toHours(seconds int) float64 {
	return float64(seconds) / 3600
}

func getTemplateData(users []User) templateData {
	TITLE := GetEnvVarWithDefault("TS_TITLE", "Watch Time Scoreboard")
	FOOTER := GetEnvVarWithDefault("TS_FOOTER", "Made with ❤️")

	sort.Slice(users, func(i, j int) bool {
		return users[i].WatchTime > users[j].WatchTime
	})

	return templateData{
		Title:  TITLE,
		Footer: FOOTER,
		Users:  users,
	}
}
