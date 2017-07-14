package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Jun-Chang/my-grpc/proto"
)

func main() {
	person := os.Getenv("PERSON")
	protocol := os.Getenv("PROTOCOL")

	var (
		reply string
		err   error
	)
	switch protocol {
	case "grpc":
		reply, err = callGRPC(context.Background(), person)
	case "rest":
		reply, err = callREST(person)
	default:
		log.Fatalf("unknown protocol %s", protocol)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

var (
	grpcConn *grpc.ClientConn
	grpcCli  proto.GreeterClient
	grpcOnce sync.Once
)

func callGRPC(ctx context.Context, personName string) (string, error) {
	grpcOnce.Do(func() {
		grpcConn, _ = grpc.Dial("localhost:13301", grpc.WithInsecure())
		grpcCli = proto.NewGreeterClient(grpcConn)
	})
	reply, err := grpcCli.Greet(ctx, &proto.Person{
		Name: personName,
	})
	if err != nil {
		return "", err
	}
	return reply.Body, nil
}

func callREST(personName string) (string, error) {
	res, err := http.Get("http://localhost:13302?person=" + personName)
	if err != nil {
		return "", err
	}
	reply := &proto.Reply{}
	d := json.NewDecoder(res.Body)
	if err := d.Decode(reply); err != nil {
		return "", err
	}
	return reply.Body, nil
}
