package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ln, e := net.Listen("tcp", ":8001")
	checkError(e)
	defer ln.Close()
	fmt.Println("Listening at localhost:8001")

	for {
		conn, e := ln.Accept()
		checkError(e)

		m, e := bufio.NewReader(conn).ReadString('\n')
		checkError(e)
		fmt.Printf("Recieving: %s\nmessage: %s", conn.RemoteAddr(), m)

		writeToBooks(m[2:])
		conn.Write([]byte("Book saved to file!\n"))
		conn.Close()
	}
}

func writeToBooks(m string) {
	books := fmt.Sprintf("%s%s", getBooks(), m)

	f, e := os.Create("books.csv")
	checkError(e)
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	_, e = w.WriteString(books)
	checkError(e)
}

func getBooks() string {
	r, e := os.ReadFile("books.csv")
	checkError(e)
	return string(r)
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
