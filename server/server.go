package main

import (
	"context"
	"fmt"
	chitchat "handin3chitchat/chitchat"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

var lamport int32

type Server struct {
	chitchat.UnimplementedChatServiceServer
}

var msgCh = make(chan *chitchat.ClientMessage)
var userCount = 0

func main() {

	lamport = 0

	port := "5678" //set the default port to 5678

	//initialize the listener on the specified port. net.Listen listens for incoming connections with tcp socket
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Could not listen at port: %s : %v", port, err)
	}
	//close listener in case of unexpected exit.
	defer listen.Close()
	log.Println("Listening at: " + port)

	//Make instance of grpc server
	grpcServer := grpc.NewServer()
	//make instance of chat server structure
	serverStructure := Server{}
	//We associate the chat service implementation, represented by the serverStructure structure
	//with the (new and empty) gRPC server.
	chitchat.RegisterChatServiceServer(grpcServer, &serverStructure)

	//grpc listen and serve
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
		log.Print(err)
	}
	select {}
}

var mutex sync.Mutex //We define a mutex to ensure only one client accesses the users map at a time
//to maintain integrity

func (s *Server) Join(User *chitchat.User, userStream chitchat.ChatService_JoinServer) error {
	fmt.Print(User.Name)
	//Compare lamport timestamps and select the highest value, then increment to maintain lamport time stamp across chat room.
	userLamport := User.Lamport
	//Use mutex to ensure consistency in the lamport timestamp across the server and all connected clients.
	mutex.Lock()
	lamport = max(lamport, userLamport)
	lamport++
	defer mutex.Unlock() //defer?

	mutex.Lock() //We lock the user map to ensure consistency in the shared resource when joining a user.
	userCount++
	defer mutex.Unlock() //defer?

	// Send and broadcast a welcome message
	welcomeMessage := fmt.Sprintf("Participant %s joined Chitty-Chat at Lamport time %d", User.Name, lamport)
	//create message
	message := &chitchat.ClientMessage{
		Name:    User.Name,
		Text:    welcomeMessage,
		Lamport: lamport,
	}
	msgCh <- message
	return nil
}

func (s *Server) SendMessage(ctx context.Context, message *chitchat.ClientMessage) (*chitchat.SentChatResponse, error) {
	fmt.Println(" - ", message.Lamport, ":", message.Text)

	messageLamport := message.Lamport
	//Use mutex to ensure consistency in the lamport timestamp across the server and all connected clients.
	//mutex.Lock()
	lamport = max(lamport, messageLamport)
	lamport++
	//defer mutex.Unlock()
	message.Lamport = lamport

	// Broadcast the message to all connected clients with the updated Lamport timestamp.
	//Broadcast(message.Name, message.Text, message.Lamport)

	//make it for range of users
	for i := 0; i < userCount; i++ {
		msgCh <- message
	}
	// Return a success response (Nothing) to indicate a successful message broadcast (in accordance with the protocol ).
	return &chitchat.SentChatResponse{}, nil
}

func BroadcastListener(User *chitchat.User) *chitchat.ClientMessage {
	return <-msgCh
}
