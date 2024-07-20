package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, e := net.Listen("tcp", ":8000")
	checkError(e)
	defer ln.Close()
	fmt.Println("Listening on port 8000")

	for {
		c, e := ln.Accept()
		checkError(e)

		r, e := bufio.NewReader(c).ReadString('\n')
		checkError(e)
		fmt.Printf("Recieving: %s\nmessage: %s", c.RemoteAddr(), r)

		m := ""
		switch r[0] {
		case 'B':
			m = sendToBooks(r)
		case 'M':
			m = sendToMovies(r)
		default:
			m = "Failed to send, bad object type!"
		}

		c.Write([]byte(fmt.Sprintf("%s\n", m)))
		c.Close()
	}
}

func sendToBooks(m string) string {
	c, e := net.Dial("tcp", "localhost:8001")
	checkError(e)
	defer c.Close()

	fmt.Fprintf(c, m+"\n")
	r, e := bufio.NewReader(c).ReadString('\n')
	checkError(e)
	return r
}

func sendToMovies(m string) string {
	c, e := net.Dial("tcp", "localhost:8002")
	checkError(e)
	defer c.Close()

	fmt.Fprintf(c, m+"\n")
	r, e := bufio.NewReader(c).ReadString('\n')
	checkError(e)
	return r
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
