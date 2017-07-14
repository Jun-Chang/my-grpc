package main

import (
	"errors"
	"fmt"

	"golang.org/x/net/context"

	"github.com/Jun-Chang/my-grpc/proto"
)

type Greeter struct{}

func (g Greeter) Greet(ctx context.Context, person *proto.Person) (*proto.Reply, error) {
	if person == nil {
		return nil, errors.New("person is nil")
	}
	return &proto.Reply{
		Body: fmt.Sprintf("Hello %s.", person.Name),
	}, nil
}
