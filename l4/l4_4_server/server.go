// curl -X POST -d '{"name": "Ivan", "age": 50}' http://localhost:8080/create

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// {
// 	"name": "Ivan",
// 	"age": 20
// }

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) toString() string {
	return fmt.Sprintf("name is %s and age is %d\n", u.Name, u.Age)
}

type service struct {
	store map[string]*User
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error() + " " + string(content)))
			return
		}

		s.store[u.Name] = &u
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + u.Name))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for _, u := range s.store {
			response += u.toString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {

	srv := service{make(map[string]*User)}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("I am alive"))
	})
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAll)

	http.ListenAndServe("localhost:8080", mux)
}
