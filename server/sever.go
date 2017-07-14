package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"google.golang.org/grpc"

	"github.com/Jun-Chang/my-grpc/proto"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		log.Println("listen grpc 13301")
		serveGRPC()
		wg.Done()
	}()
	go func() {
		log.Println("listen rest 13302")
		serveREST()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done")
}

func serveGRPC() {
	l, _ := net.Listen("tcp", ":13301")
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, Greeter{})

	s.Serve(l)
}

func serveREST() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Query()["person"]
		if len(p) == 0 {
			fmt.Fprint(w, `{"body": "Tell me your name"}`)
			return
		}
		fmt.Fprintf(w, `{"body": "Hello %s."}`, p[0])
	})

	http.ListenAndServe(":13302", nil)
}
