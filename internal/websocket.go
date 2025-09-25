package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteJSON(PlayerCounts)
			if err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}
}

func InitWebSocket() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("Server ws:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
