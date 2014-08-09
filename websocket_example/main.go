package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",
    func (w http.ResponseWriter, req *http.Request) {
        s := websocket.Server{Handler: websocket.Handler(Echo)}
        s.ServeHTTP(w, req)
    });

	if err := http.ListenAndServe(":1800", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
