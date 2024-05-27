package main

import (
	"net/http"
)

const (
	Port = "8080"
)

type Message struct {
	Content string
	Author  string
}

func main() {
	messages := []Message{{"hi", "test1"}, {"hello", "test2"}, {"yo", "test1"}}

	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.PostFormValue("name")
			if name == "" {
				http.Error(w, "name is empty", http.StatusBadRequest)
				return
			}
			http.SetCookie(w, &http.Cookie{Name: "name", Value: name, HttpOnly: true})
			http.Redirect(w, r, "/chat", http.StatusFound)
			return
		}

		_, err := r.Cookie("name")
		if err == nil {
			http.Redirect(w, r, "/chat", http.StatusFound)
			return
		}

		err = nameSelectionTemplate.Execute(w, messages)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		name, err := r.Cookie("name")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		type Data = struct {
			Messages []Message
			Name     string
		}
		data := Data{messages, name.Value}
		err = chatTemplate.Execute(w, data)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			author, err := r.Cookie("name")
			if err != nil {
				http.Error(w, "name cookie not set", http.StatusBadRequest)
				return
			}

			message := r.PostFormValue("message")
			if message == "" {
				http.Error(w, "message is empty", http.StatusBadRequest)
				return
			}

			messages = append(messages, Message{message, author.Value})
		}

		author, err := r.Cookie("name")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		type Data = struct {
			Messages []Message
			Name     string
		}
		data := Data{messages, author.Value}
		err = chatTemplate.ExecuteTemplate(w, "messages", data)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":"+Port, nil)
}
