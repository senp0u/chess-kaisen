package main

import (
    "html/template"
    "net/http"
	"log"
    "time"
    "os"
    "github.com/gorilla/websocket"
)

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
    tmpl, err := template.ParseFiles("templates/welcome.html", "templates/board.html")
    if err != nil{
        panic(err)
    }
    tmpl.Execute(w, games)
}

func test(conn *websocket.Conn, ch <-chan string) {
    defer conn.Close()
    for {
        time.Sleep(5 * time.Second) 
        board, err := os.ReadFile("templates/board.html")
        if err != nil{
            panic(err)
        }
        if err := conn.WriteMessage(1, board); err != nil {
            log.Println("Error: ", err)
            break
        }
    }
}

func wsPlay(w http.ResponseWriter, r *http.Request) {
    // Upgrade upgrades the HTTP connection to WebSocket
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
        return
	}

	ch := make(chan string)

	go test(conn, ch)
}