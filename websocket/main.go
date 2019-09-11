package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	go func() {
		http.HandleFunc("/", IndexHandler)
		http.HandleFunc("/ws", WebSocketHandler)
		http.ListenAndServe("localhost:8080", nil)
	}()

	done := make(chan os.Signal, 1)
	// ...
	done <- os.Interrupt
	// ...

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	select {
	case <-interrupt:
	case <-done:
	}

	log.Println("shutting down")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexHtml, _ := os.Open("websocket.html")
	indexData, _ := ioutil.ReadAll(indexHtml)
	fmt.Fprint(w, string(indexData))
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := websocket.Upgrade(w, r, nil, 1024, 1024)
	mType, bts, _ := conn.ReadMessage()
	fmt.Println(string(bts))
	conn.WriteMessage(mType+1, []byte("message from back-end"))
	conn.Close()
}
