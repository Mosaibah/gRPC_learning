package chat

import (
	"context"
	"log"
)
	

type Server struct {
	ChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server yay!"}, nil
}

func (s *Server) Feeling(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("haay this from client: %s", in.Body)
	return &Message{Body: "Im feeling so good alhamdullah"}, nil
}
