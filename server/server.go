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
	userStreams map[int32]*UserStream
	chitchat.UnimplementedChatServiceServer
}

// create map of all userstreams (streans to connected clients)
var userStreams = make(map[int32]*UserStream)
var lamport int32

func main() {

	//initialize lamport clock
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

	//Compare lamport timestamps and select the highest value, then increment to maintain lamport time stamp across chat room.
	userLamport := User.Lamport
	//Use mutex to ensure consistency in the lamport timestamp across the server and all connected clients.
	lamport = max(lamport, userLamport)
	lamport++

	// Send and broadcast a welcome message
	welcomeMessage := fmt.Sprintf("Participant %s joined Chitty-Chat at Lamport time %d", User.Name, lamport)
	message := &chitchat.ClientMessage{
		Name:    User.Name,
		Text:    welcomeMessage,
		Lamport: lamport,
	}
	s.BroadcastChatMessage(context.Background(), message)

	//Add user to map of userstreams.
	newUserStream := &UserStream{
		UserId: User.Id,    // Set the user's ID
		Stream: userStream, // Set the gRPC stream
	}

	userStreams[User.Id] = newUserStream

	select {}
}

func (s *Server) BroadcastChatMessage(ctx context.Context, message *chitchat.ClientMessage) (*chitchat.Confirmation, error) {
	fmt.Println(" - ", message.Lamport, ":", message.Text)

	messageLamport := message.Lamport
	//Use mutex to ensure consistency in the lamport timestamp across the server and all connected clients.
	//mutex.Lock()
	lamport = max(lamport, messageLamport)
	lamport++
	//defer mutex.Unlock()
	message.Lamport = lamport

	//Send the message to all connected users by cycling through the userStreams and sending the message.
	for _, userStream := range userStreams {
		if err := userStream.Stream.Send((*chitchat.ServerMessage)(message)); err != nil {
			log.Fatalf("Failed to send send message to client with id %d", userStream.UserId)
		}
	}

	return &chitchat.Confirmation{}, nil
}

//not used anymore
/*func (s *Server) BroadcastListener(context.Context, *chitchat.User) (*chitchat.ClientMessage, error) {
	for {
		clientMessage := <-msgCh
		// Send the received message to the client
		return clientMessage, nil
	}
}*/
