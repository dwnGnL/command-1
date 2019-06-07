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
	var logined bool
	readedPassHash := readPasswordHash()
	connections = append(connections,conn)
	var userName string
	
	for !logined {
		userName,_=bufio.NewReader(conn).ReadString('\n')
		userName = userName[:len(userName)-2]
		pass,_:=bufio.NewReader(conn).ReadString('\n')
		pass= pass[:len(pass)-2]
	if userName == "root" && checkHash(pass, readedPassHash) {
		logined = true
		fmt.Println("Hello, root")
		_,err:=conn.Write([]byte("true;"))
		if err != nil{
			fmt.Println(err)
		}
	} else {
		_,err:=conn.Write([]byte("false;"))
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("Invalid credentials")
	}
}
	_,err := conn.Write([]byte("Welcome to chat Mr(s) "+ userName + "\n;"))

	if err != nil{
		fmt.Println(err)
	}

	for !exit {
		reader := bufio.NewReader(conn)
		command, err := reader.ReadString(';')
		if err !=nil{
			fmt.Println(err.Error())
			conn.Close()
			removeConn(conn)
			fmt.Print(userName+" is Offline")
			break
		}
		text:=allMake(command[:len(command)-1])
		_,err2 := conn.Write([]byte(text))

			if err2 !=nil{
				fmt.Println(err.Error())
				conn.Close()
				removeConn(conn)
				fmt.Print(userName+" is Offline")
				break
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

func removeConn(conn net.Conn){
	var i int

	for i = range connections{
		if connections[i] == conn {
			break
		}
	}


	fmt.Println(i)

	if len(connections) > 1{
		connections = append(connections[:i],connections[i+1:]...)
	}else{
		connections = nil
	}

}

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