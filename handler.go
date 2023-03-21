package main

import (
	"context"
	pb "counterGrpc"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var mutex = &sync.Mutex{}
var dic = make(map[string]chan int32)
var tokenCh = make(chan string)

const (
	address ="localhost:50051"
)



func main() {
	go grpcHitCount()

	http.HandleFunc("/grpcCount", grpcHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func grpcHandler(w http.ResponseWriter, r *http.Request){
	//create token
	token := generateToken()

	//create channel for goroutine to return
	hit := make(chan int32)
	//NEED: close the channel you don't need
	defer close(hit)

	mutex.Lock()
	//write to dic <token, return channel>
	dic[token] = hit
	mutex.Unlock()
	log.Printf("ch created in handler: %v", hit)

	//pass token to goroutine listening in main
	tokenCh <- token
	log.Printf("send token %v to channel", token)

	counterGrpc := <- hit
	log.Printf("get counter %v back from ch", counterGrpc)
	fmt.Fprint(w, "The hit of this page from grpcCount is ", counterGrpc)
	//NEED: delete dic that returned already

	mutex.Lock()
	delete(dic, token)
	mutex.Unlock()
}

func grpcHitCount(){

	//set up a connection to server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCounterQueueClient(conn)

	for {
		//get token by global channel
		token := <-tokenCh
		log.Printf("get token %v by channel", token)

		log.Printf("sending token %v", token)

		//Contact the server and print out its response
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		
		defer cancel()
		r, err := c.RequestCounter(ctx, &pb.CounterGrpc{Token: token, Count: 0})
		if err != nil{
			log.Fatalf("cound not request counter: %v", err)
		}

		//get respond from grpc, get counter
		log.Printf("get counter respond: token: %v, counter: %v", r.Token, r.Count)
		myCount := r.Count
		myToken := r.Token

		mutex.Lock()
		ch := dic[myToken]
		mutex.Unlock()
		//send counter back to specific channel
		ch <- myCount
	}
}

func generateToken() string{
	b := uuid.New()
	return  b.String()
}
