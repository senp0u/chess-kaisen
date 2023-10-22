package main

import (
    "html/template"
    "net/http"
	"log"
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

func play(w http.ResponseWriter, r *http.Request) {
    username := "Test"//r.PostFormValue("username")
    if username == "" {
        w.Header().Set("x-missing-field", "username")
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if games.isGameFull(){
        panic("The game is full")
    }

    // Upgrade upgrades the HTTP connection to WebSocket
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
		w.Header().Set("x-missing-field", "username")
        w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
        return
	}
    defer conn.Close()

	chn := make(chan string)

	go func() {
		defer close(chn)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Message:", message)
				log.Println("Error:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	/*for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}


    
    /*games.addPlayerToGame(username)*/
    tmpl, err := template.ParseFiles("templates/welcome.html", "templates/board.html")
    if err != nil{
        panic(err)
    }
    tmpl.Execute(w, games)
}