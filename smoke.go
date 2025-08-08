package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6359")
	if err != nil {
		log.Fatal("error occured when listening to TCP socket: ", err)
	}

	fmt.Println("Started TCP server at port :6359")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected:", conn.RemoteAddr())
			} else {
				fmt.Println("Read error:", err)
			}

			return
		}

		msg := string(buf[:n])
		fmt.Println(msg)
	}
}
