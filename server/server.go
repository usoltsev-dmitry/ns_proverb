package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const address = "0.0.0.0:12345"
const protocol = "tcp4"

// Go-Поговорки
var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func main() {
	// Запуск сетевой службы по протоколу TCP
	listener, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

// Обработчик каждые 3 сек отдает клиенту случайную поговорку. Соединение с клиентом не прерывается.
func handleConn(conn net.Conn) {
	for {
		proverb := getRandomProverb()
		_, err := conn.Write([]byte(proverb + "\n"))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		time.Sleep(3 * time.Second)
	}
}

// Функция отдает случайную поговорку из среза proverbs
func getRandomProverb() string {
	rand.Seed(time.Now().UnixNano())
	return proverbs[rand.Intn(len(proverbs))]
}
