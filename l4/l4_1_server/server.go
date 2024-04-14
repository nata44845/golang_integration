// Клиент-серверное приложение, сервер, возвращает клиенту текст в верхнем регистре
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("server is running")
	con, err := lis.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		line, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("line: ", string(line))

		upperline := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upperline)); err != nil {
			log.Fatalln(err)
		}
	}
}
