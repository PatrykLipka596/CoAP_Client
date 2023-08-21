package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/plgd-dev/go-coap/v3/udp"
)

func main() {
	co, err := udp.Dial("192.168.0.153:5688")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	path := "/a"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := co.Get(ctx, path)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	log.Printf("Response payload: %v", resp.String())
	body, err := resp.ReadBody()
	if err != nil {
		log.Fatalf("Couldn't read body: %v", err)
	}
	log.Printf("Response body: %v", string(body[:]))
}
