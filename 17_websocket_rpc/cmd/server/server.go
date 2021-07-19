package main

import (
	"go_course_thinknetika/17_websocket_rpc/pkg/api"
	"go_course_thinknetika/17_websocket_rpc/pkg/chat"
)

type Service struct {
	api  *api.API
	chat *chat.Chat
}

func main() {
	service := new()
	go service.publishMessages()
	service.api.Run()
}

func new() *Service {
	chat := chat.New()

	s := Service{
		api:  api.New(":8000", chat),
		chat: chat,
	}

	return &s
}

func (s *Service) publishMessages() {
	for msg := range s.chat.MsgQueue {
		for _, c := range s.chat.Clients {
			c <- msg
		}
	}
}
