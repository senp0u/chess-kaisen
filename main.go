package main

import (
    "html/template"
    "net/http"
    "log"
)

func main() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        tmpl, err := template.ParseFiles("templates/index.html", "templates/username-form.html")
        if err != nil {
            panic(err)
        }
        tmpl.Execute(w, nil)
    })

    http.HandleFunc("/username-form/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/username-form.html"))
        tmpl.Execute(w, nil)
    })

    http.HandleFunc("/play/", func(w http.ResponseWriter, r *http.Request) {
        username := r.PostFormValue("username")
        if username == "" {
            w.Header().Set("x-missing-field", "username")
		    w.WriteHeader(http.StatusBadRequest)
	        return
        }
        tmpl, err := template.ParseFiles("templates/welcome.html", "templates/board.html")
        if err != nil{
            panic(err)
        }
        tmpl.Execute(w, struct{ Username string }{username})
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}