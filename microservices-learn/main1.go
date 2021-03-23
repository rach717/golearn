package main

import (
	"log"
	"handlers"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.Lstdflags)
	hh := handlers.NewHello(l)
	http.ListenAndServe(":9090", nil)
}
