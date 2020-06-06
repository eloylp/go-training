package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	go func() {
		addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
		if err != nil {
			log.Fatal(err)
		}
		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			for {
				b := make([]byte, 1024)
				_, err = c.Read(b)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("server received: ", string(b))
				_, err = c.Write([]byte("pong"))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	addr, err := net.ResolveTCPAddr("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for {
		if _, err := c.Write([]byte("ping")); err != nil {
			log.Fatal(err)
		}
		b := make([]byte, 1024)
		_, _ = c.Read(b)
		fmt.Println(string(b))
		time.Sleep(time.Second)
	}
}
