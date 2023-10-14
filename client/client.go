package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Подключение к сетевой службе.
	conn, err := net.Dial("tcp4", "localhost:12345")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Буфер для чтения данных из соединения.
	reader := bufio.NewReader(conn)

	// Считывание массива байт до перевода строки.
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		// Обработка ответа.
		fmt.Println(string(b))
	}
}
