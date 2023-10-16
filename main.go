package main

import (
    "html/template"
    "net/http"
    "log"
    "fmt"
)

type Player struct{
    Username string
    Color string
}

type Game struct{
    White Player
    Black Player
}
func (g Game) startGame(){

    return 
}

func (g Game) isGameFull() bool{
    return g.White.Username != "" && g.Black.Username != ""
}

func (g *Game) addPlayerToGame(username string){
    if g.White.Username == ""{
        g.White = Player{
            Username: username,
            Color: "White",
        }
    }else{
        g.Black = Player{
            Username: username,
            Color: "Black",
        }
        g.startGame()
    }
}

var games Game

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

    log.Fatal(http.ListenAndServe(":8080", nil))
}


func play(w http.ResponseWriter, r *http.Request) {
    username := r.PostFormValue("username")
    if username == "" {
        w.Header().Set("x-missing-field", "username")
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if games.isGameFull(){
        panic("The game is full")
    }
    
    games.addPlayerToGame(username)
    fmt.Println(games.White)
    tmpl, err := template.ParseFiles("templates/welcome.html", "templates/board.html")
    if err != nil{
        panic(err)
    }
    tmpl.Execute(w, games)
}