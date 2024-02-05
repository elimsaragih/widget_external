/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	widget "google.golang.org/grpc/examples/helloworld/contract"
	pb "github.com/elimsaragih/widget_external/grpc_proto"
)

var (
	port      = flag.Int("port", 50051, "The server port")
	byteArray = []byte{123, 34, 70, 105, 114, 115, 116, 78, 97, 109, 101, 34, 58, 34, 75, 114, 105, 115, 104, 110, 97, 34, 44, 34, 76, 97, 115, 116, 78, 97, 109, 101, 34, 58, 34, 71, 117, 114, 114, 97, 109, 34, 44, 34, 73, 68, 34, 58, 49, 50, 51, 125}
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// integrate widget
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	data := widget.Widget{
		Title:    "server 1",
		SubTitle: in.Name,
	}
	temp, _ := json.Marshal(data)
	return &pb.HelloReply{Message: "Hello again " + in.GetName(), Body: temp}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
