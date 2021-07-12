package api

import (
	"fmt"
	"go_course_thinknetika/17_websocket_rpc/pkg/chat"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type API struct {
	port   string
	router *mux.Router
	chat   *chat.Chat
}

func New(port string, chat *chat.Chat) *API {
	s := API{
		port:   port,
		router: mux.NewRouter(),
		chat:   chat,
	}
	return &s
}

func (api *API) Run() {
	api.endpoints()
	http.ListenAndServe(api.port, api.router)
}

func (api *API) endpoints() {
	api.router.HandleFunc("/send", api.handleSendMessage).Methods(http.MethodGet)
	api.router.HandleFunc("/messages", api.handleMessages).Methods(http.MethodGet)
}

func (api *API) handleSendMessage(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	log.Println("Client connected")

	err = ws.WriteMessage(websocket.TextMessage, []byte("Enter the password!\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	auth := false
	for {
		mt, r, err := ws.ReadMessage()
		if err != nil {
			log.Println("1", err)
			break
		}
		resp := string(r)
		if !auth && resp != "password" {
			log.Println("invalid password")
			err = ws.WriteMessage(mt, []byte("Invalid password, try again!\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if auth == true {
			log.Println("Message:", resp)

			api.chat.Broadcast(resp)

			err = ws.WriteMessage(mt, []byte("Sent\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if string(resp) == "password" {
			auth = true
			err = ws.WriteMessage(mt, []byte("Ok\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func (api *API) handleMessages(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	strID, client := api.chat.Subscribe()
	defer api.chat.Unsubscribe(strID)
	defer ws.Close()

	for msg := range client {
		err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
