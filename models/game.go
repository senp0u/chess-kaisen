package model

type Game struct{
    White Player
    Black Player
    Ch chan string
    Board [64]Piece
}

func (g Game) StartGame(){
    //ToDo
    return 
}


func (g *Game) RemovePlayer(){
    g.White.Username = ""
}

func (g Game) IsGameFull() bool{
    return g.White.Username != "" && g.Black.Username != ""
}

func (g *Game) AddPlayerToGame(username string){
    if g.White.Username == ""{
        g.White = Player{
            Username: username,
            Color: "White",
        }
        g.Ch = make(chan string)
    }else{
        g.Black = Player{
            Username: username,
            Color: "Black",
        }
        g.Ch <- username
        g.StartGame()
    }
}
