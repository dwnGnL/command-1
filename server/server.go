package main

import(
	"fmt"
	"bufio"
	"net"
	// "os"
)

var connections []net.Conn

func main()  {
	
	ln,err :=net.Listen("tcp",":8111")

	if err !=nil{
		panic(err)
	}

	for{
		conn,err:=ln.Accept()

		if err != nil{
			fmt.Println(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn){
	var exit bool
	connections = append(connections,conn)
	userName,_:=bufio.NewReader(conn).ReadString('\n')
	userName = userName[:len(userName)-2]
	_,err := conn.Write([]byte("Welcome to chat Mr(s) "+ userName + "\n"))

	if err != nil{
		fmt.Println(err)
	}

	for !exit {
		reader := bufio.NewReader(conn)
		command, _ := reader.ReadString(';')
		text:=allMake(command[:len(command)-1])
		_,err := conn.Write([]byte(text))

			if err !=nil{
				fmt.Println(err.Error())
			}

		
	}

	// for{
	// 	text, err := bufio.NewReader(conn).ReadString('\n')

	// 	if err != nil{
	// 		conn.Close()
	// 		removeConn(conn)
	// 		broadCastMsg(userName + " is offline \n",conn)
	// 		break
	// 	}

	// 	broadCastMsg(userName+":"+text,conn)
		
	// }
}

// func removeConn(conn net.Conn){
// 	var i int

// 	for i = range connections{
// 		if connections[i] == conn {
// 			break
// 		}
// 	}

// 	fmt.Println(i)

// 	if len(connections) > 1{
// 		connections = append(connections[:i],connections[i+1:]...)
// 	}else{
// 		connections = nil
// 	}

// }

// func broadCastMsg(msg string,sourceConn net.Conn){
// 	for _,conn:=range connections{
// 		if sourceConn != conn{
// 			_,err := conn.Write([]byte(msg))

// 			if err !=nil{
// 				fmt.Println(err.Error())
// 			}
// 		}
// 	}
// 	msg = msg[:len(msg)-1]
// 	fmt.Println(msg)
// }