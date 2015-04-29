package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	addr string
	port int
)

func init() {
	flag.StringVar(&addr, "addr", "", "listen to address")
	flag.IntVar(&port, "port", 3000, "listen to port")
}
func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), nil)
}
