package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Утилита telnet

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Требования:
	1. Программа должна подключаться к указанному хосту (ip или
доменное имя + порт) по протоколу TCP. После подключения
STDIN программы должен записываться в сокет, а данные
полученные и сокета должны выводиться в STDOUT
	2. Опционально в программу можно передать таймаут на
подключение к серверу (через аргумент --timeout, по
умолчанию 10s)
	3. При нажатии Ctrl+D программа должна закрывать сокет и
завершаться. Если сокет закрывается со стороны сервера,
программа должна также завершаться. При подключении к
несуществующему сервер, программа должна завершаться
через timeout
*/

var (
	timeoutFlag = flag.Int("timeout", 10, "Таймаут на подключение к серверу. По-умолчанию 10s")
)

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("there are no args")
	} else if flag.NArg() == 1 {
		log.Fatal("there are no port in args")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	conn, err := net.DialTimeout("tcp",
		fmt.Sprintf("%s:%s", host, port),
		time.Duration(*timeoutFlag)*time.Second)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(conn)

	go copyTo(os.Stdout, conn)
	go copyTo(conn, os.Stdin)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT)

	<-quit
}
