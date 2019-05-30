package main

import (
	"bufio"
	"fmt"
	"os"
)

var ID int
var db = map[string]*[]Student{}
	var emptySlice = []Student{}

func main() {
	var login string
	var password string
	var logined bool
	var exit bool

	readedPassHash := readPasswordHash()

	for !logined {
		fmt.Println("Enter username: ")
		fmt.Scan(&login)
		fmt.Println("Enter pass:")
		fmt.Scan(&password)

		if login == "root" && checkHash(password, readedPassHash) {
			logined = true
			fmt.Println("Hello, root")
		} else {
			fmt.Println("Invalid credentials")
		}
	}

	for !exit {
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		text:=allMake(command)

		fmt.Println(text)
	}

}
