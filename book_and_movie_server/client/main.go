package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for {
		res, msg := "", ""
		fmt.Println(strings.Join(clientPrompt(), "\n"))
		fmt.Scanln(&res)
		switch res {
		case "1":
			msg = fmt.Sprintf("B,%s\n", getDetails("book"))
		case "2":
			msg = fmt.Sprintf("M,%s\n", getDetails("movie"))
		default:
			os.Exit(0)
		}
		sendMessage(msg)
	}
}

func clientPrompt() []string {
	return []string{
		"**************************************",
		"Place your order by selecting a number",
		"**************************************",
		"1. Purchase book",
		"2. Purchase movie",
		"Press any other key to exit",
		"**************************************",
	}
}

func getDetails(t string) string {
	fmt.Printf("Enter %s details\n", t)
	return fmt.Sprintf("%s,%s,%s", getPrompt(t, "name"), getPrompt(t, "quantity"), getPrompt(t, "price"))
}

func getPrompt(t string, p string) string {
	out := ""
	fmt.Printf("Enter %s %s: ", t, p)
	fmt.Scanln(&out)
	return out
}

func sendMessage(m string) {
	c, e := net.Dial("tcp", "localhost:8000")
	checkError(e)
	defer c.Close()

	fmt.Fprintf(c, m+"\n")
	msg, e := bufio.NewReader(c).ReadString('\n')
	checkError(e)
	fmt.Printf("Message from server: %s", msg)
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
