package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/elimsaragih/widget_external/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	widget "google.golang.org/grpc/examples/helloworld/contract"
)

const (
	defaultName = "world"
)

var (
	addr     []*string
	name     = flag.String("Perso", defaultName, "Widget Perso")
	register = []string{"server1", "server2"}
)

func main() {

	temp1 := flag.String("addr", "localhost:50051", "the address to connect to")
	temp2 := flag.String("addr2", "localhost:50052", "the address to connect to")
	flag.Parse()
	addr = append(addr, temp1)
	addr = append(addr, temp2)

	c := make(map[string]pb.GreeterClient, 0)
	// Set up a connection to the server.
	for i, v := range addr {
		conn, err := grpc.Dial(*v, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}

		defer conn.Close()
		c[register[i]] = pb.NewGreeterClient(conn)
	}

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ini ada di repo untuk get widget dari grpc
	r, err := c["server2"].SayHelloAgain(ctx, &pb.HelloRequest{Name: "pers1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())

	// cek dia pakai widget apa, biar kita tau harus di umarshal ke widget mana
	var emp widget.Widget
	err = json.Unmarshal(r.GetBody(), &emp)

	if err != nil {
		fmt.Println("Can;t unmarshal the byte array")
		return
	}

	emp.GetWidget()
}
