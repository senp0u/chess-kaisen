package main

import (
    "html/template"
    "net/http"
    "log"
    "github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}


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

    http.HandleFunc("/play/", play)

    http.HandleFunc("/wsplay/", wsPlay)

    log.Fatal(http.ListenAndServe(":8080", nil))
}