package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := "message from handler"
	if _, err := w.Write([]byte(text)); err != nil {
		log.Fatalln(err)
	}
}

func main() {

}
