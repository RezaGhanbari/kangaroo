package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	pb "kangaroo/consignment-service/proto/consignment"
	"log"
	"os"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, err
	}
	return consignment, err
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("and error occured: %v", err)
		}
		return
	}()
	client := pb.NewShippingServiceClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
