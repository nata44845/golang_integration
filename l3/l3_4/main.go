// Мьютекс
package main

import (
	"sync"
)

type server struct {
	status string
	sync.Mutex
}

func (s *server) Alive() {
	s.Lock()
	s.status = "alive"
	s.Unlock()
}

func (s *server) Down() {
	s.Lock()
	s.status = "down"
	s.Unlock()
}

func main() {
	s := server{}
	for i := 0; i < 1000; i++ {
		go s.Alive()
		go s.Down()
	}

}
