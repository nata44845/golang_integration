package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const proxyAddr string = "localhost:9000"

var (
	counter int = 0
)

func setCounter() {
	if counter == 0 {
		counter = 1
	} else {
		counter = 0
	}
}

var instances = map[int]string{0: "http://localhost:8080", 1: "http://localhost:8081"}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	setCounter()

	resp, err := http.Get(instances[counter] + "/users")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	textBytes, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(fmt.Sprintf("instance %d\n", counter)))
	w.Write(textBytes)
}

func friends(w http.ResponseWriter, r *http.Request) {
	setCounter()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := http.Get(instances[counter] + fmt.Sprintf("/friends/%d", id))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	textBytes, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(fmt.Sprintf("instance %d\n", counter)))
	w.Write(textBytes)
}

func user(w http.ResponseWriter, r *http.Request) {
	setCounter()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := http.Get(instances[counter] + fmt.Sprintf("/user/%d", id))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	textBytes, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(fmt.Sprintf("instance %d\n", counter)))
	w.Write(textBytes)
}

func create(w http.ResponseWriter, r *http.Request) {
	setCounter()

	textProxy, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(textProxy)
	resp, err := http.Post(instances[counter]+"/user/create", "application/json", bytes.NewBuffer([]byte(text)))

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	textBytes, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(fmt.Sprintf("instance %d\n", counter)))
	w.Write(textBytes)
}

func delete(w http.ResponseWriter, r *http.Request) {
	setCounter()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", instances[counter]+fmt.Sprintf("/user/%d", id), nil)

	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	textBytes, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(fmt.Sprintf("instance %d\n", counter)))
	w.Write(textBytes)
}

func setAge(w http.ResponseWriter, r *http.Request) {
	setCounter()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	textProxy, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(textProxy)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", instances[counter]+fmt.Sprintf("/user/%d", id), bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	textBytes, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(fmt.Sprintf("instance %d\n", counter)))
	w.Write(textBytes)
}

func makeFriends(w http.ResponseWriter, r *http.Request) {
	setCounter()

	textProxy, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(textProxy)
	resp, err := http.Post(instances[counter]+"/make_friends", "application/json", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	textBytes, err := ioutil.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(fmt.Sprintf("instance %d\n", counter)))
	w.Write(textBytes)
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Сервер запущен"))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Сервер запущен"))
	})

	r.Get("/test", test)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", users)
	})

	//1,3
	r.Route("/user", func(r chi.Router) {
		r.Post("/create", create)
		r.Get("/{id}", user)
		r.Delete("/{id}", delete)
		r.Put("/{id}", setAge)
	})

	//2
	r.Route("/make_friends", func(r chi.Router) {
		r.Post("/", makeFriends)
	})

	//4
	r.Route("/friends", func(r chi.Router) {
		r.Get("/{id}", friends)
	})

	log.Fatalln(http.ListenAndServe(proxyAddr, r))
}
