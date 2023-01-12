package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/middlewere"
)

type User struct {
	Username   string `json:"username"`
	Fullname   string `json:"fullname"`
	ProfileUrl string `json:"profileUrl"`
}

var allUsers = []User{
	{
		Username:   "user1",
		Fullname:   "fullname1",
		ProfileUrl: "http1",
	},
	{
		Username:   "user2",
		Fullname:   "fullname2",
		ProfileUrl: "http2",
	},
	{
		Username:   "user3",
		Fullname:   "fullname3",
		ProfileUrl: "http3",
	},
	{
		Username:   "user4",
		Fullname:   "fullname4",
		ProfileUrl: "http4",
	},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("user handler ", r.Method, r.URL)
	if r.Method == http.MethodGet {
		q := r.URL.Query()
		f := q.Get("filter")

		if f == "" {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(allUsers)
			return
		}

		var users []User
		for _, u := range allUsers {
			if strings.Contains(u.Username, f) {
				users = append(users, u)
			}
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)

	} else {
		msg := http.StatusText(http.StatusMethodNotAllowed)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}

func main() {
	e := echo.New()
	e.Use(middlewere.Logger())

	e.GET("/users", usersHandler)

	port := "8080"
	log.Println("starting port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
