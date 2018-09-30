package main

import (
	"encoding/json"
	"flag"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type hashed struct {
	Number string `json:"number"`
	Hash   string `json:"hash"`
}

func init() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error loading .env file")
	}
}



func main() {
	//Redis connection
	cli, err := redis.Dial("tcp",":"+os.Getenv("REDIS_PORT"))
	if err != nil {
		// handle error
		panic(err)
	}
	defer cli.Close()
	flag.Parse()
	args := flag.Args()
	if len(args[0]) < 6 {
		println("first arg must me more then 6 chars")
		os.Exit(1)
	}

	times, _ := strconv.Atoi(args[1])
	if times <= 0 {
		println("second arg must me more then 1")
		os.Exit(2)
	}
	sendToRedis(args[0], times,cli)
}

// SendToRedis function sends number and hash to redis N times
func sendToRedis(num string, times int,cli redis.Conn) {
	//generating hash N times
	for i:=0;i<times;i++{
		HSum := generateHash(num)
	h := hashed{
		Number: num,
		Hash:   string(HSum),
	}
	json, _ := json.Marshal(h)
	cli.Do("LPUSH", os.Getenv("LIST"), json)
	//after push generate new number and repeat
	num = regenerateNumber(num)
	}
}
// regenerateNumber regenerates given number
// by deleting last 4 symbol to random
func regenerateNumber(num string) string {
		num = num[:len(num)-4]
		z := rand.Intn(9999)
		//get random more than 4 symbols number
		for z < 1000 {
			z = rand.Intn(9999)
		}
		num += strconv.Itoa(z)
	return num
}

// regenerateHash regenerates given hash with sha3 algorithm
func generateHash(num string) []byte {
	hash := sha3.New512()
	hash.Write([]byte(num))
	HSum := hash.Sum(nil)
	return HSum
}
