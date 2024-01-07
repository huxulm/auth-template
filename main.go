package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/huxulm/auth-template/routers"
	"github.com/quasoft/memstore"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "host:port")
}

func main() {
	mux := http.NewServeMux()
	routers.Register(mux)

	store := memstore.NewMemStore(
		[]byte("authkey123"),
		[]byte("enckey12341234567890123456789012"),
	)
	store.Options.MaxAge = 3600
	routers.SetStore(store)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.Serve(ln, mux); err != nil {
		log.Fatal(err)
	}
}
