package main

import (
	"context"
	pb "counterGrpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

var counter int32


/*type CounterQueueImpl struct {
}

func NewCounterQueueImpl() *CounterQueueImpl{
	return &CounterQueueImpl{}
}*/

type server struct {
	pb.UnimplementedCounterQueueServer
}

func (s *server) RequestCounter(ctx context.Context, in *pb.CounterGrpc) (*pb.CounterGrpc, error){
	log.Printf("Server counterApp recieved: token: %v, count: %v", in.GetToken(), in.GetCount())
	counter = counter + 1
	return &pb.CounterGrpc{Token: in.GetToken(), Count: counter}, nil
}

func main(){
	counter = 0
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("fail to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCounterQueueServer(s, &server{})
	if err:= s.Serve(lis); err != nil{
		log.Fatalf("fail to serve: %v", err)
	}

}

/*func counterApp(){

	counter := 0

	//counterCh := make(chan counterRabbitMQ)

	for {
		////Counter stay local in counterApp now so this
		//dequeue grpc
		rabbitMQRecieve(*ch, sendQname, counterCh)
		countMQ := <-counterCh
		log.Printf("in app, recieve token and counter %v", countMQ)
		counter = counter + 1
		countMQ.Count = counter

		//send respond to grpc
		err = ch.Publish(
			"",        // exchange
			receiveQname, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        json,
			})
		failOnError(err, "Failed to publish a message")
		log.Printf("in app, send %v to rabbitMQ", countMQ)
	}
}*/
