// go run server.go
// git bash
// curl -X POST -d '{"name":"Ilya","age":28,"friends":[]}' http://localhost:8080/user/create
// curl -X POST -d '{"name":"Maxim","age":25,"friends":[]}' http://localhost:8080/user/create
// curl -X GET http://localhost:8080/users
// curl -X POST -d '{"source_id":"1","target_id":"2"}' http://localhost:8080/make_friends
// curl -X GET http://localhost:8080/users
// curl -X GET http://localhost:8080/user/2
// curl -X GET http://localhost:8080/friends/2
// curl -X GET http://localhost:8080/friends/1
// curl -X GET http://localhost:8080/friends/3
// curl -X DELETE http://localhost:8080/user/1
// curl -X GET http://localhost:8080/users
// curl -X PUT -d '{"new age":"20"}' http://localhost:8080/user/2

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"

	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// {
// 	"name": "Ivan",
// 	"age": 20
//  "friends": []
// }

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

const fileName string = "users.txt"

func max(users map[int]*User) int {
	var maxNumber int
	for maxNumber = range users {
		break
	}
	for n := range users {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

type User struct {
	Name    string  `json: "name"`
	Age     int     `json: "age"`
	Friends []*User `json:"friends"`
}

func (u *User) toString() string {
	response := fmt.Sprintf("%s возраст %d\n", u.Name, u.Age)
	if len(u.Friends) > 0 {
		response += "Друзья:\n"
		for _, f := range u.Friends {
			response += f.Name + "\n"
		}
	}
	return response
}

type service struct {
	counter int
	store   map[int]*User
}

func (s *service) Users(w http.ResponseWriter, r *http.Request) {
	s.LoadDataFromFile()
	response := ""
	for i, u := range s.store {
		response += fmt.Sprintf("Пользователь %d: ", i) + u.toString()
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (s *service) User(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	s.LoadDataFromFile()
	u := s.store[id]
	if u != nil {
		response := u.toString()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Пользователь не найден %d", id)))
	}
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
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

	s.LoadDataFromFile()
	s.counter++
	s.store[s.counter] = &u

	//Запись в файл json
	err = s.SaveDataToFile()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Пользователь создан %d %s", s.counter, u.Name)))
}

func (s *service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var objmap map[string]json.RawMessage
	if err := json.Unmarshal(content, &objmap); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + " " + string(content)))
		return
	}

	var str string
	if err := json.Unmarshal(objmap["source_id"], &str); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + " " + string(content)))
		return
	}

	id1, err := strconv.Atoi(str)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + " " + string(content)))
		return
	}

	if err := json.Unmarshal(objmap["target_id"], &str); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + " " + string(content)))
		return
	}

	id2, err := strconv.Atoi(str)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + " " + string(content)))
		return
	}

	u1 := s.store[id1]
	u2 := s.store[id2]
	//Проверка на зацикливание
	if u1 != nil && u2 != nil {

		idx := slices.Contains(u2.Friends, u1)
		if !idx {
			s.LoadDataFromFile()
			u1.Friends = append(u1.Friends, u2)
			//Запись в файл json
			err = s.SaveDataToFile()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%s и %s теперь друзья", u1.Name, u2.Name)))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Добавление приведет к зацикливанию списка")))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Пользователь не найден " + string(content)))
	}
}

func (s *service) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	s.LoadDataFromFile()
	u := s.store[id]
	if u != nil {
		for _, user := range s.store {
			friends := []*User{}
			for _, friend := range user.Friends {
				if friend != u {
					friends = append(friends, friend)
				}
			}
			user.Friends = friends
		}
		delete(s.store, id)
		//Запись в файл json
		err = s.SaveDataToFile()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Пользователь удален " + u.Name))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Пользователь не найден %d", id)))
	}
}

func (s *service) Friends(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	s.LoadDataFromFile()
	u := s.store[id]
	if u != nil {
		response := fmt.Sprintf("Друзья пользователя %s: \n", u.Name)
		for _, f := range u.Friends {
			response += f.Name
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Пользователь не найден %d", id)))
	}
}

func (s *service) setAge(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	s.LoadDataFromFile()
	u := s.store[id]
	if u != nil {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var objmap map[string]json.RawMessage
		if err := json.Unmarshal(content, &objmap); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error() + " " + string(content)))
			return
		}

		var str string

		if err := json.Unmarshal(objmap["new age"], &str); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error() + " " + string(content)))
			return
		}

		age, err := strconv.Atoi(str)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error() + " " + string(content)))
			return
		}

		u.Age = age

		//Запись в файл json
		err = s.SaveDataToFile()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Возраст изменен пользователь %s возраст %d", u.Name, age)))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Пользователь не найден %d", id)))
	}
}

func (s *service) LoadDataFromFile() {
	file, err := os.Open(fileName)
	if err != nil {
		file, _ = os.Create(fileName)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err := json.Unmarshal(content, &s.store); err != nil {
		fmt.Println(err.Error())
		for k := range s.store {
			delete(s.store, k)
		}
		s.counter = 0
	} else {
		s.counter = max(s.store)
	}
}

func (s *service) SaveDataToFile() error {
	var b bytes.Buffer
	userJson, _ := json.Marshal(s.store)
	b.WriteString(string(userJson))
	err := ioutil.WriteFile(fileName, b.Bytes(), 0777)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func main() {

	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	srv := service{0, make(map[int]*User)}

	//Загрузка из файла
	srv.LoadDataFromFile()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Сервер запущен"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", srv.Users)
	})

	//1,3
	r.Route("/user", func(r chi.Router) {
		r.Post("/create", srv.Create)
		r.Get("/{id}", srv.User)
		r.Delete("/{id}", srv.Delete)
		r.Put("/{id}", srv.setAge)
	})

	//2
	r.Route("/make_friends", func(r chi.Router) {
		r.Post("/", srv.MakeFriends)
	})

	//4
	r.Route("/friends", func(r chi.Router) {
		r.Get("/{id}", srv.Friends)
	})

	log.Fatalln(http.ListenAndServe("localhost:"+port, r))
}
