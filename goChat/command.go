//处理用户的命令，如更换房间，获取用户信息等功能

package goChat

import (
	"fmt"
	"net"
	//"goChat"
	"strconv"

)

//命令的结构体
type Command struct {
	Conn *net.TCPConn
	Cmd string
	CInfo *ClientsInfo
	RoomList []Room
	Clt *Client
	
}

//处理命令的入口
func (c Command) DisCmd(cmd string) bool{
	
	switch(cmd) {
		case "showclients":
			fmt.Println("show client")
			c.Cmd = "showclients"
			c.showOnlineClient()
			return true
		case "showroom":
			fmt.Println("show room")
			c.Cmd = "showroom"
			c.showRoom()	
			return true
		case "createroom":
			fmt.Println("create room")
			c.Cmd = "createroom"
			c.createRoom()
			return true	
		case "changeroom2":
			fmt.Println("change room 2")
			c.Cmd = "changeroom"
			c.changeRoom(1)
			return true	
		case "changeroom3":
			fmt.Println("change room 3")
			c.Cmd = "changeroom"
			c.changeRoom(2)
			return true	
		case "roominfo":
			fmt.Println("room info")
			c.Cmd = "roominfo"
			c.roominfo()
			return true	
		case "exit":
			fmt.Println("exit !!")
			c.Cmd = "exit"
			c.exit()
			return true	
		default:
			return false
			//fmt.Println("command not found!")	
	}
}

//显示当前在线用户
func (c Command) showOnlineClient() {
	//c.Conn.Write([]byte("showOnli neClient:\n"))
	c.Conn.Write([]byte("NAME          IP          ROOM"))
	for _, v := range c.CInfo.Clients {
		s := v.Name + "      " + v.IP + "       " + strconv.Itoa(v.Rom.Num)
		fmt.Println(s)
		c.Conn.Write([]byte(s))
	}
}

//显示房间
func (c Command) showRoom() {
	//c.Conn.Write([]byte("show room"))
	//roomList := *c.RoomList
	c.Conn.Write([]byte("ROOM_NAME      ROOM_NUM\n"))
	for _, v := range c.RoomList {
		s := v.Name + "         " + strconv.Itoa(v.Num)
		c.Conn.Write([]byte(s))
	}
} 

//更换房间
func (c Command) changeRoom(num int) {
	c.Clt.Rom.RemoveClient(*c.Clt)
	c.RoomList[num].AddClient(*c.Clt)
	c.Clt.Rom = c.RoomList[num]
}

//新建房间
func (c Command) createRoom() {
	//i := len(c.RoomList) 
//	s := c.Clt.Name + "建立的聊天室"
//	r := NewRoom(3, s)
//	c.RoomList = append(c.RoomList, r)
//	go r.Start()
	
	
}

//删除房间
func (c Command) removeRoom() {
	
}

//显示当前房间信息
func (c Command)roominfo() {
	roomNum := strconv.Itoa(c.Clt.Rom.Num)  //房间号码
	roomName := c.Clt.Rom.Name  //房间名
	
	i := 0   //房间人数
	var s string 
	for k, _ := range c.Clt.Rom.Client {
		i++
		s = s + " " + k
	}
	roominfo := []byte("当前房间号码为 " + roomNum + " 房间名是 " + roomName)
	roomclient := []byte("当前房间人数为 " + strconv.Itoa(i) + " 他们是 " + s)
	if i !=0 {
		c.Conn.Write(roominfo)
		c.Conn.Write(roomclient)
	}
	
}

//用户正常退出
func (c Command)exit() {
	c.Clt.Rom.RemoveClient(*c.Clt)
	delete(c.CInfo.Clients, c.Clt.Name)
}