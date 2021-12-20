package game_pub_server

import (
	"flag"
	"fmt"
	"github.com/wansnow/calculation_server/config"
	"github.com/wansnow/calculation_server/server/calculation_server/service/hub_service"
	"log"
	"net/http"
)

var addr = flag.String("addr", fmt.Sprintf(":%d", config.ServerC.StartGameServerPort), "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func Run() {
	flag.Parse()
	hub := hub_service.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub_service.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
