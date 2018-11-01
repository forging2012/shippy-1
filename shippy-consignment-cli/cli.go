package main

import (
	"encoding/json"
	pb "github.com/CcccFz/shippy/shippy-consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
)

const (
	defaultFilename = "shippy-consignment-cli/consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiZmNlMmU0OGYtMTA0ZS00ODkxLWE4YzMtOWYxNGVkYjBhMzEwIiwibmFtZSI6IkNjY2NGeiIsImNvbXBhbnkiOiJHb2xkZW4gUmlkZ2UiLCJlbWFpbCI6ImNjY2NmekAxNjMuY29tIiwicGFzc3dvcmQiOiIkMmEkMTAkcC9LTXhVbG11NmU0aW5odU84TlFCdWgxdTVNLmp4Llo0OWNCODlPRWxWUmhOZUxtL1F3b1cifSwiZXhwIjoxNTAwMCwiaXNzIjoiZ28ubWljcm8uc3J2LnVzZXIifQ.Zn0uicdc4HMM0otfNKSKCcpJsvL3kJC4viEyRDFF2hY"

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
