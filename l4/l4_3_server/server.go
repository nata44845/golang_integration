//http server получает данные через пробел
// curl -X POST -d "name 24" http://localhost:8080/create

package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type service struct {
	store map[string]string
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		splitedContent := strings.Split(string(content), " ")
		s.store[splitedContent[0]] = string(content)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + splitedContent[0]))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for _, user := range s.store {
			response += user + "\n"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {

	srv := service{make(map[string]string)}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("I am alive"))
	})
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAll)

	http.ListenAndServe("localhost:8080", mux)
}
