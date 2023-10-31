package main

import (
	"context"
	"fmt"
	chitchat "homework3/chitchat"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type UserStream struct {
	UserId int32
	Stream chitchat.ChatService_JoinServer // The gRPC stream
}

type Server struct {
	chitchat.UnimplementedChatServiceServer
}

// create map of all userstreams (streans to connected clients)
var userStreams = make(map[int32]*UserStream)
var lamport int32

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

	//We make an instance of grpc server and chat server structure
	grpcServer := grpc.NewServer()
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

var mutex sync.Mutex

func (s *Server) Broadcast(ctx context.Context, message *chitchat.ClientMessage) (*chitchat.Confirmation, error) {
	messageLamport := message.Lamport
	//Use mutex to ensure consistency in the lamport timestamp across the server and all connected clients.
	mutex.Lock()
	lamport = max(lamport, messageLamport)
	lamport++
	mutex.Unlock()

	message.Lamport = lamport

	fmt.Println(" - ", message.Lamport, message.Name, ":", message.Text)

	//Send the message to all connected users by cycling through the userStreams and sending the message.
	for _, userStream := range userStreams {
		if err := userStream.Stream.Send((*chitchat.ServerMessage)(message)); err != nil {
			log.Fatalf("Failed to send message to client with id %d", userStream.UserId)
		}
	}

	return &chitchat.Confirmation{}, nil
}

func (s *Server) Join(User *chitchat.User, userStream chitchat.ChatService_JoinServer) error {

	//Compare lamport timestamps and select the highest value, then increment to maintain lamport time stamp across chat room.
	userLamport := User.Lamport
	mutex.Lock()
	lamport = max(lamport, userLamport)
	lamport++
	mutex.Unlock()

	// Send and broadcast a welcome message
	welcomeMessage := fmt.Sprintf("Participant %s joined Chitty-Chat at Lamport time %d", User.Name, lamport)
	message := &chitchat.ClientMessage{
		Name:    "SERVER MESSAGE",
		Text:    welcomeMessage,
		Lamport: lamport,
	}
	s.Broadcast(context.Background(), message)
	s.Broadcast(context.Background(), message)

	//Add user to map of userstreams.
	newUserStream := &UserStream{
		UserId: User.Id,
		Stream: userStream,
	}
	//Use mutex to ensure consistency in shared resource userStreams.
	mutex.Lock()
	userStreams[User.Id] = newUserStream
	mutex.Unlock()
	//keep method running to keep the userstream open.
	select {}
}

func (s *Server) Leave(ctx context.Context, User *chitchat.User) (*chitchat.Confirmation, error) {

	userLamport := User.Lamport
	//Use mutex to ensure consistency in the lamport timestamp across the server and all connected clients.
	mutex.Lock()
	lamport = max(lamport, userLamport)
	lamport++
	mutex.Unlock()

	//delete the userstream mapped to the given id from the userstreams map.
	//Use mutex to ensure consistency in shared resource userStreams.
	mutex.Lock()
	delete(userStreams, User.Id)
	mutex.Unlock()

	//Broadcast leave message
	leaveMessage := fmt.Sprintf("Participant %s left Chitty-Chat at Lamport time %d", User.Name, lamport)
	message := &chitchat.ClientMessage{
		Name:    "SERVER MESSAGE",
		Text:    leaveMessage,
		Lamport: lamport,
	}
	s.Broadcast(context.Background(), message)
	return &chitchat.Confirmation{}, nil
}
