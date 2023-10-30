package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"
	"unicode/utf8"

	"homework3/chitchat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type chatClientStruct struct {
	stream chitchat.ChatService_JoinClient
	id     int32
	name   string
}

var lamport int32
var user *chitchat.User

func main() {

	//Server address where gRPC server is running
	const serverAddress = "localhost:5678"

	//We initialize lamport clock
	lamport = 0

	//we create insecure transport credentials (in the context of this assignment we choose not to worry about security):
	transportCreds := insecure.NewCredentials()
	//Establish a grpc connection to the server using addres and tansport credentials
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(transportCreds))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server ... : %v\n", err)
	}
	//defer call to conn.Close() to ensure connection is closed when main method exits.
	defer conn.Close()

	//create client and user
	client := chitchat.NewChatServiceClient(conn)
	chatClient := chatClientStruct{}
	user = chatClient.CreateUser(client)

	//sleep to simulate wait time for connection to be established...
	log.Println("Connecting to the gRPC server at ... : " + serverAddress)
	time.Sleep(time.Millisecond * time.Duration(1000))

	//Initialize a join stream and set the join stream in chatClient
	joinStream, err := client.Join(context.Background(), user)
	if err != nil {
		log.Fatalf("Ouch. Failed to join the chat: %v\n", err)
	}
	chatClient.stream = joinStream

	//print welcome message.
	log.Printf("\n\nHello, %s. \nYou can disconnect with '/disconnect' \n\nWrite a message ...\n", user.Name)

	//We start go routines for sending and recieving messages.
	go chatClient.SendChatMessage(client)
	go chatClient.ReceiveMessage(client, user)

	//create channel for listening for client closing down unexpectedly (for instance ctrl+c)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		//Disconnect the user and print "Disconnected", then exit.
		log.Println("Disconnected")
		client.Leave(context.Background(), user)
		os.Exit(0)
	}()

	//keep the main function running
	select {}
}

func (chatClient *chatClientStruct) SendChatMessage(client chitchat.ChatServiceClient) {
	for {
		//read user message from the console and decide what to do
		message, err := readUserInput()
		if utf8.RuneCountInString(message) > 128 {
			log.Println("Your message must be no longer than 128 characters!")
		} else if err != nil {
			log.Fatalf("Ouch. Failed to read your chat message from the console: %v ", err)
		} else if message == "/disconnect" {
			//increment lamport and call Leave method to disconnect.
			lamport++
			client.Leave(context.Background(), user)
			//since the user won't recieve the broadcast leave message from the server after disconnecting we print a leave message for the client.
			log.Print("You have left the chat!")
			os.Exit(0)
		} else {
			lamport++
			//Create new clientMessage and send it to the server by calling BroadcastChatmessage.
			clientMessage := &chitchat.ClientMessage{
				Name:    chatClient.name,
				Text:    message,
				Lamport: lamport,
			}
			_, err2 := client.BroadcastChatMessage(context.Background(), clientMessage)
			if err2 != nil {
				log.Fatalf("Failed to send the clientMessage to server: %v\n", err2)
			}
		}
	}
}

func (chatClient *chatClientStruct) ReceiveMessage(client chitchat.ChatServiceClient, user *chitchat.User) {
	for {

		//recieve a message from the server
		userStreamServerMessage, err := chatClient.stream.Recv()
		if err != nil {
			log.Fatalf("Failed to recieve message from server: %v\n", err)
		}
		if userStreamServerMessage == nil {
			log.Fatalf("serverMessage returned nil")
		}

		//Find lamport timestamp of incoming message, select the highest and increment.
		incomingLamport := userStreamServerMessage.Lamport
		lamport = max(lamport, incomingLamport)
		lamport++

		//Displaying the recieved chat message with lamport time stamp:
		log.Printf(" - [%d] %s: %s", lamport, userStreamServerMessage.Name, userStreamServerMessage.Text)
	}
}

func (chatClient *chatClientStruct) CreateUser(client chitchat.ChatServiceClient) *chitchat.User {
	//Generate a random id:
	randSrc := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSrc)
	id := int32(randGen.Intn(999999)) // Generate a random integer and cast to int32
	chatClient.id = id

	//Ask client for username:
	for {
		fmt.Println("Please enter your username and press 'enter'!")
		username, err := readUserInput()
		if err != nil {
			log.Fatalf("Failed to read username: %v", err)
			continue //prompt the user to enter username againc if username not accepted
		}
		chatClient.name = username
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
	//Trim message spaces from beginning and end.
	userInput = strings.TrimSpace(userInput)
	return userInput, err
}
