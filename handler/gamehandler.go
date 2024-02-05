package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gorilla/websocket"
	model "github.com/senp0u/chess-kaisen/models"
	view "github.com/senp0u/chess-kaisen/views"
)


var game model.Game

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func Run(){
    http.Handle("/", templ.Handler(view.Index()))

    http.HandleFunc("/cancel/", func (w http.ResponseWriter, r *http.Request) {
        game.RemovePlayer()
        view.UsernameForm().Render(r.Context(),w)
    })

    http.HandleFunc("/play/", play)

    http.HandleFunc("/wsplay/", wsPlay)

    log.Fatal(http.ListenAndServe(":8080", nil))
}


func play(w http.ResponseWriter, r *http.Request) {
    username := r.PostFormValue("username")
    if username == "" {
        w.Header().Set("x-missing-field", "username")
        w.WriteHeader(http.StatusBadRequest)
        view.UsernameForm().Render(r.Context(), w)
    }
    if game.IsGameFull(){
        panic("The game is full")
    }
    game.AddPlayerToGame(username)
    view.Game(game).Render(r.Context(), w)
}

func PlayGame(conn *websocket.Conn, r *http.Request) {
    defer conn.Close()
    fmt.Println("In Websocket")
    if !game.IsGameFull(){
        game.Black.Username = <-game.Ch
    }
    for {
        fmt.Println("In for")

        time.Sleep(5 * time.Second)
        component := view.Board(game.Board)
        componentBytes, err := ComponentToBytes(r.Context(), &component)
        if err != nil{
            log.Println("Error generating []bytes from teml.Component")
        }
        if err := conn.WriteMessage(1, componentBytes); err != nil {
            log.Println("Error: ", err)
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
    go PlayGame(conn, r)
}

func ComponentToBytes(ctx context.Context, c *templ.Component) (b []byte, err error) {
	buffer := templ.GetBuffer()
	defer templ.ReleaseBuffer(buffer)
	if err = (*c).Render(ctx, buffer); err != nil {
	    return []byte{}, err
	}
	return buffer.Bytes(), nil
}
