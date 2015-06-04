package main 

import (
	"net"
	"fmt"
	"os"
	"os/exec"
)

var Conn *net.TCPConn

func main() {
	
	//用户名称，用户聊天的昵称
	var username string
	//生成ipaddr地址，其中后面的127。。。为服务器地址和监听端口
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9091")
	checkErr(err)
	//建立连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	Conn = conn
	//checkErr(err)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		//连接成功后会首先要求输入用户名称
		fmt.Print("请输入用户名:")
		fmt.Scan(&username)
		conn.Write([]byte(username))
		checkErr(err)
		//建立连接
		createChat(conn, username)
	}
	
	
	//err = conn.CloseWrite()
	
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func createChat(conn *net.TCPConn, username string) {
//	sendMsg := make([]byte, 1024)
//	recvMsg := make([]byte, 1024)
//	for {
//		fmt.Printf("%s :", username)
//		_, err := fmt.Scan(&sendMsg)
//		checkErr(err)
//		
//		conn.Write(sendMsg)
//		
//		c, err := conn.Read(recvMsg)
//		if err != nil {
//			fmt.Println("消息未正确送达：", err.Error())
//		}
//		fmt.Println(string(recvMsg[0:c]))
//	}
	go sendMsg(conn, username)
	receiveMsg(conn, username)
}
//向服务器发送数据
func sendMsg(conn *net.TCPConn, username string) {
	sendMsg := make([]byte, 1024)
	for {
		//fmt.Printf("%s :", username)
		_, err := fmt.Scan(&sendMsg)
		checkErr(err)
		if command(string(sendMsg)) == false {
			conn.Write(sendMsg)	
		}
		
	}
}

//从服务器获取数据
func receiveMsg(conn *net.TCPConn, username string){
	recvMsg := make([]byte, 1024)
	for {
		c, err := conn.Read(recvMsg)
		if err != nil {
			fmt.Println("消息未正确送达：", err.Error())
		}
		fmt.Println(string(recvMsg[0:c]))
	}
}

//服务端的一些控制命令
func command(cmd string) bool{
	switch(cmd) {
		
		case "clear":  //清屏
			fmt.Println("clear screen")
			clear()
			return true
		case "exit":  //退出程序
			fmt.Println("exit")
			exit()
			return true		
		default:
			return false		
	} 
}

//清屏
func clear() {
	cmd := exec.Command("/bin/sh", "-c" , "clear")
	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}
}
func exit() {
	os.Exit(1)
	Conn.Write([]byte("exit"))
}