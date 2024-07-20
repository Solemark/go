package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ln, e := net.Listen("tcp", ":8002")
	checkError(e)
	defer ln.Close()
	fmt.Println("Listening at localhost:8002")

	for {
		conn, e := ln.Accept()
		checkError(e)

		m, e := bufio.NewReader(conn).ReadString('\n')
		checkError(e)
		fmt.Printf("Recieving: %s\nmessage: %s", conn.RemoteAddr(), m)

		writeToMovies(m[2:])
		conn.Write([]byte("Movie saved to file!\n"))
		conn.Close()
	}
}

func writeToMovies(m string) {
	movies := fmt.Sprintf("%s%s", getMovies(), m)

	f, e := os.Create("movies.csv")
	checkError(e)
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	_, e = w.WriteString(movies)
	checkError(e)
}

func getMovies() string {
	r, e := os.ReadFile("movies.csv")
	checkError(e)
	return string(r)
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
