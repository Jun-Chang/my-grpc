package main

import (
	"testing"

	"golang.org/x/net/context"
)

var (
	person = "Bob"
	ctx    = context.Background()
)

func BenchmarkGRPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := callGRPC(ctx, person); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkREST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := callREST(person); err != nil {
			b.Fatal(err)
		}
	}
}
