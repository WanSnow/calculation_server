package hub_service

import (
	"encoding/json"
	"github.com/wansnow/calculation_server/model/game"
	"time"
)

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			go pubGameResult(client.gameId, h.broadcast, client.end)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.end <- 1
				close(client.send)
				close(client.end)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				g := new(game.Game)
				err := json.Unmarshal(message, g)
				if err != nil {
					break
				}
				if g.GameId != client.gameId {
					continue
				}
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func pubGameResult(gameId string, b chan []byte, end chan int) {
	gu := game.NewUse()
loop:
	for {
		getGame, err := gu.GetGame(gameId)
		if err != nil {
			return
		}
		select {
		case <-end:
			break loop
		default:
			marshal, err := json.Marshal(getGame)
			if err != nil {
				return
			}
			b <- marshal
		}
		time.Sleep(100 * time.Millisecond)
	}

}
