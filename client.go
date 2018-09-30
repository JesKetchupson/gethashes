package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dialer := websocket.Dialer{}

for{
	//Establishing connection
	con,_,err:= dialer.Dial("ws://"+os.Getenv("HOST")+":"+os.Getenv("PORT")+"/ws",nil)

	if err!=nil{
		panic(err)
	}
	//Reading it
	_,Mbody,_:= con.ReadMessage()
	fmt.Printf("%s\n",Mbody)
}
}
