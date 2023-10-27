package main

import (
	"bufio"
	"context"
	"fmt"
	"handin3chitchat/chitchat"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type chatClientStruct struct {
	stream chitchat.ChatService_JoinClient
	id     int32
	name   string
}

var lamport int32

func main() {

	//Server address
	const serverAddress = "localhost:5678"
	log.Println("Client connecting to ... : " + serverAddress)

	//We initialize lamport clock
	lamport = 0

	//we create insecure transport credentials (in the context of this assignment we choose not to worry about security):
	transportCreds := insecure.NewCredentials()
	//Establish a grpc connection to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(transportCreds))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server ... : %v\n", err)
	}
	defer conn.Close()

	client := chitchat.NewChatServiceClient(conn)
	chatClient := chatClientStruct{}

	//create user
	user := chatClient.CreateUser(client)

	//Initialize a join stream and set the join stream in chatClient
	joinStream, err := client.Join(context.Background(), user)
	if err != nil {
		log.Fatalf("Ouch. Failed to join the chat: %v\n", err)
	}
	chatClient.stream = joinStream

	//We start go routines for sending and recieving messages.
	go chatClient.SendChatMessage(client)
	go chatClient.ReceiveMessage()

	//keep the main function running
	select {}
}

func (chatClient *chatClientStruct) SendChatMessage(client chitchat.ChatServiceClient) {
	for {
		//Promt user to write messages in chat
		fmt.Println("Send a message: ")

		//read user message from the console.
		message, err := readUserInput()
		if err != nil {
			log.Fatalf("Ouch. Failed to read your chat message from the console: %v ", err)
		}
		//We increment lamport in order to give message a lamport timestamp.
		lamport++
		//Create new clientMessage and send it to the server.
		clientMessage := &chitchat.ClientMessage{
			Name:    chatClient.name,
			Text:    message,
			Lamport: lamport,
		}
		_, err = client.SendMessage(context.Background(), clientMessage)
		if err != nil {
			log.Fatalf("Failed to send the clientMessage to server: %v\n", err)
		}
	}
}

func (chatClient *chatClientStruct) ReceiveMessage() {
	for {
		//recieve a message from the server
		serverMessage, err := chatClient.stream.Recv()
		if err != nil {
			log.Fatalf("Failed to recieve message from server: %v\n", err)
		}

		//Find lamport timestamp of incoming message, select the highest and increment.
		incomingLamport := serverMessage.Lamport
		lamport = max(lamport, incomingLamport)
		lamport++

		//Displaying the recieved chat message with lamport time stamp:
		log.Printf("[%d] %s : \n %s", lamport, serverMessage.Name, serverMessage.Text)
	}
}

func (chatClient *chatClientStruct) CreateUser(client chitchat.ChatServiceClient) *chitchat.User {
	//Generate a random id:
	randSrc := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSrc)
	id := int32(randGen.Intn(10000)) // Generate a random integer and cast to int32
	chatClient.id = id

	//Ask client for username:
	for {
		fmt.Println("Please enter your username and press 'enter'!")
		username, err := readUserInput()
		if err != nil {
			log.Fatalf("Failed to read username: %v", err)
			continue //prompt the user to enter username again
		}
		chatClient.name = username
		fmt.Println("Hello, %s. You are new ready to start chatting!")
		//break out of the loop since we have a username
		break
	}

	//Generate new user:
	var user = &chitchat.User{
		Id:      chatClient.id,
		Name:    chatClient.name,
		Lamport: lamport,
	}
	return user
}
func readUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	//Trim message whitespace from beginning and end.
	userInput = strings.TrimSpace(userInput)
	return userInput, err

}