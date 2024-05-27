package main

import (
	"chat/api/chat"
	"chat/api/messages"
	"chat/api/root"
	"chat/templates"
	"net/http"
)

const (
	Port = "8080"
)

func main() {
	http.Handle("/static/", http.FileServer(http.FS(templates.StaticFiles)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			root.Get(w, r)
		case http.MethodPost:
			root.Post(w, r)
		}
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		chat.Get(w, r)
	})

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			messages.Get(w, r)
		case http.MethodPost:
			messages.Post(w, r)
		}
	})

	http.ListenAndServe(":"+Port, nil)
}
