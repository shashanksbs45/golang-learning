package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("opening connection")
	c, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading to webSocket %s", err)
		return
	}
	defer func() {
		log.Println("closing connection")
		c.Close()
	}()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error while reading the message from client: %s", err)
			return
		}
		log.Printf("Receive message %s", string(message))
		if strings.Trim(string(message), "\n") != "start" {
			err = c.WriteMessage(websocket.TextMessage, []byte("You did not say the magic word!"))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
				return
			}
			continue
		}
		log.Println("Start responsding to client")
		for {
			err := c.WriteMessage(websocket.TextMessage, []byte("response"))
			if err != nil {
				log.Printf("Error %s when sending message to client", err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}

}

func main() {
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}
	http.Handle("/", webSocketHandler)
	log.Print("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
