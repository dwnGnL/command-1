package main

import(
	"fmt"
	"net"
	"log"
	"os"
	"io"
	"bufio"
)

func main()  {
	conn,err:=net.Dial("tcp",":8111")
	var logined bool
	if err!=nil{
		log.Println("нет соединения ",err.Error())
		return
	}

	
	for !logined {
	fmt.Println("Enter username: ")
	name,err:= bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println("Enter pass:")
	pass,err:= bufio.NewReader(os.Stdin).ReadString('\n')
	
	
	if err!=nil{
		log.Println("не можем отправить имя на сервер",err.Error())
		return
	}
	conn.Write([]byte(name))
	conn.Write([]byte(pass))
	text,_:= bufio.NewReader(conn).ReadString(';')
	if text == "true;"{
		logined=true
	}
	}
	
	go ReadMsg(conn)
	WriteMsg(conn)

}

func WriteMsg(conn net.Conn){
	for{
		text,err:= bufio.NewReader(os.Stdin).ReadString(';')

		if err!=nil{
			log.Println("не можем отправить сообщение на сервер",err.Error())
			panic(err)
		}
		conn.Write([]byte(text))
	}
	
}

func ReadMsg(conn net.Conn)  {
	for {
		text,err:= bufio.NewReader(conn).ReadString(';')
		
		if err == io.EOF{
			log.Println("Server don`t work")
			conn.Close()
			panic(err)
		}else if err!=nil{
			log.Println("Server don`t work")
			conn.Close()
			panic(err)
		}
		fmt.Println(text[:len(text)-1])
	}
}