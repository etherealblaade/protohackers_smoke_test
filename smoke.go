package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6359")
	if err != nil {
		log.Fatal("listen:", err)
	}
	log.Println("TCP echo on :6359")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}
		go func(c net.Conn) {
			defer c.Close()
			_, _ = io.Copy(c, c)
		}(conn)
	}
}
