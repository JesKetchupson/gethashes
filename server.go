package main

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)
func init() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws", Ws).Methods("GET")
	println("server started at "+ os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}

var upgrader = websocket.Upgrader{}
//slice of websockets connection
var conChan []*websocket.Conn

// Ws handler to interact thru websockets
func Ws(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	//adding new connection on slice
	conChan = append(conChan, conn)

	cli, _ := redis.Dial("tcp", ":"+os.Getenv("REDIS_PORT"))
	func(conn *websocket.Conn) {
		for {
			re, _ := cli.Do("MONITOR")
			//Finding json string in redis monitor output
			var validID = regexp.MustCompile(`{([\s\S]+?)}`)
			match := validID.FindStringSubmatch(re.(string))
			if match != nil {
				for _, b := range conChan {
					b.WriteMessage(1, []byte(match[0]))
					time.Sleep(time.Second * 3)
				}

			}
		}
	}(conn)
}
