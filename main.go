package main

import (
    "html/template"
    "net/http"
    "log"
)

func main() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, nil)
    })
    http.HandleFunc("/play/", func(w http.ResponseWriter, r *http.Request) {
        username := r.PostFormValue("username")
        if username == "" {
            w.Header().Set("x-missing-field", "username")
		    w.WriteHeader(http.StatusBadRequest)
	        return
        }
        log.Print(username)
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}