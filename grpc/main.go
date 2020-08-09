package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const Port = ":8080"

type MessageServer struct{}

func (MessageServer) SayIt(ctx context.Context, r *Request) (*Response, error) {
	fmt.Println("Requests text:" + r.Text)
	fmt.Println("Requests sub-text:" + r.Subtext)
	response := &Response{Text: "response text", Subtext: "Got it !"}
	return response, nil
}

func main() {
	time.AfterFunc(5*time.Second, func() {
		launchClient()
	})
	server := grpc.NewServer()
	var messageServer MessageServer
	RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("serving requests")
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func launchClient() {
	conn, err := grpc.Dial(Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := NewMessageServiceClient(conn)
	r, err := SayIt(context.TODO(), client, "doing grpc !!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Text)
}

func SayIt(ctx context.Context, m MessageServiceClient, text string) (*Response, error) {
	req := &Request{Text: text, Subtext: "New message !"}
	resp, err := m.SayIt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
