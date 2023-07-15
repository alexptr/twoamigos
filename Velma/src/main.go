package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/sayHello/:message", sayHello)
	router.Run(":10000")
	//http.HandleFunc("/sayhello/", sayhello)
	//log.Fatal(http.ListenAndServe(":10000", nil))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func sayHello(c *gin.Context) {

	response := fmt.Sprintf("We are using GIN. Hello Mr. %s. This is the message: %s", randStringRunes(10), c.Param("message"))
	c.IndentedJSON(http.StatusOK, response)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	paramValue := r.URL.Path[len("/sayhello/"):]
	json.NewEncoder(w).Encode(fmt.Sprintf("Hello Mr. %s. This is your message: %s", randStringRunes(10), paramValue))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
