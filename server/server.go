package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Server() error {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return fmt.Errorf("could not allocate new listener %v", err)
	}

	for {
		errChan := make(chan error)
		msgChan := make(chan string)
		conn, err := lis.Accept()
		if err != nil {
			return fmt.Errorf("could not accept connection %v", err)
		}

		go getMessage(conn, errChan, msgChan)

		select {
		case msg := <-msgChan:
			fmt.Fprint(conn, msg)
		case err := <-errChan:
			return fmt.Errorf("error occurred %v", err)
		}
	}
}

func getMessage(conn net.Conn, errChan chan error, msgChan chan string) {
	recv, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("error receiving message", err)
		errChan <- err
		return
	}

	msgChan <- recv
}

func main() {
	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error,1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := Server(); err != nil{
			errChan <-  fmt.Errorf("error starting server %v", err)
		}
	}()

	select {
	case <-sigChan:
		fmt.Println("server stopped!")
		return
	case err := <- errChan:
		log.Fatal(err)
	}
}