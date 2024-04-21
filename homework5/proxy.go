package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter        int    = 0
	firstInstance  string = "http://localhost:8080"
	secondInstance string = "http://localhost:8081"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
	textBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	text := string(textBytes)
	if counter == 0 {
		resp, err := http.Post(firstInstance, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatalln(err)
		}

		counter++
		textBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Println(string(textBytes))
	} else {
		resp, err := http.Post(secondInstance, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatalln(err)
		}

		counter--
		textBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Println(string(textBytes))
	}

}

func main() {
	http.HandleFunc("/", handleProxy)
	log.Fatalln(http.ListenAndServe(proxyAddr, nil))
}
