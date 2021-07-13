package chat

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Chat's setup structure.
type Chat struct {
	mux      *sync.Mutex
	Clients  map[string]chan string
	MsgQueue chan string
}

// New creates Chat object.
func New() *Chat {
	c := Chat{
		mux:      &sync.Mutex{},
		Clients:  make(map[string]chan string, 0),
		MsgQueue: make(chan string),
	}

	return &c
}

// Subscribe makes a connection of Client to Server.
func (c *Chat) Subscribe() (string, chan string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	length := 32
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	strID := b.String()
	client := make(chan string)
	c.Clients[strID] = client

	return strID, client
}

// Subscribe destroys a connection from Client to Server.
func (c *Chat) Unsubscribe(cID string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.Clients, cID)
}

// Broadcast is a messages transporter.
func (c *Chat) Broadcast(message string) {
	c.MsgQueue <- message
}
