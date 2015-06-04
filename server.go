/*
	服务器端主程序
*/

package main 

import (
	"fmt"
	"net"
	"goChat"
)
//cli用于保存连接到后台的客户端信息
var Cli goChat.ClientsInfo

//创建Room列表
var RoomList []goChat.Room

//新用户登陆，注册用户列表
func registe(conn *net.TCPConn) (goChat.Client, error){
//	clientName := make([]byte, 1024) 
//	conn.Read(clientName)

//新用户注册到用户列表中
	clientName := make([]byte, 1024)
	conn.Read(clientName)
	client := goChat.NewClient(string(clientName), conn, RoomList[0])
	err := Cli.AddClient(string(clientName), client)
	if err != nil {
		fmt.Println("注册失败") 
		return client,err
	}else {
		fmt.Println("注册成功")
	}
//	for k, v := range Cli.Clients {
//		fmt.Println(k, v)
//	}
	
	//将新用户放入默认的聊天室
	RoomList[0].AddClient(client)  
	//fmt.Println(RoomList[0])
	return client,err
}


func serverChat(client goChat.Client) {
	conn := client.Conn
	msg := make([]byte, 1024) 
	command := goChat.Command{conn, "", &Cli, RoomList, &client}
	defer delete(Cli.Clients, client.Name)
	defer client.Rom.RemoveClient(client)
	for {
		
		len, err := conn.Read(msg)
		if err != nil {
			fmt.Println("客户端异常，无法与客户端连接：", err.Error())
			break
		}
		//对用户信息进行判断，看是否是命令
		if command.DisCmd(string(msg[:len])) {
			fmt.Println(command.Cmd)			
		}else {
			//fmt.Print(client.Name + ": " + string(msg[:len]))
			fmt.Println(string(msg[:len]))
			client.SendMsg(msg[:len])
			//conn.Write([]byte("信息接受成功！"))
		}
		//ch <- msg
		
		
	}
}

func main() {
	
	//生成TCPAddr格式的IP地址和端口
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":9091")
	checkErr(err)
	
	//建立listener,监听来自指定端口的请求
	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	
	//初始化cli
	Cli = goChat.ClientsInfo{make(map[string]goChat.Client)}
	
	//建立默认聊天室,并初始化	
	RoomList = []goChat.Room{goChat.NewRoom(1, "大厅聊天室"),goChat.NewRoom(2,"二号聊天室"),goChat.NewRoom(3,"三号聊天室")}
	go RoomList[0].Start()
	go RoomList[1].Start()
	go RoomList[2].Start()
	
	for {
		conn, err := ln.AcceptTCP()
		checkErr(err)	
		defer conn.Close()
		client, err := registe(conn)
		if err != nil {
			fmt.Println(err)
			break
		} else{
			//启动新的线程，为当前用户服务
		go	serverChat(client)
		} 
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

