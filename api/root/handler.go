package root

import (
	"chat/templates"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("name")
	if err == nil {
		http.Redirect(w, r, "/chat", http.StatusFound)
		return
	}

	err = templates.NameSelectionTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	if name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "name", Value: name, HttpOnly: true})
	http.Redirect(w, r, "/chat", http.StatusFound)
}
