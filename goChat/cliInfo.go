//保存客户端登陆的信息

package goChat 

import (
	"net"
	"strings"
	"fmt"
)

//var Cli ClientsInfo


//保存当前连接到客户端的用户信息

type ClientsInfo struct {
	Clients map[string]Client
}

type Client struct {
	Name  string
	IP    string
	Port  string
	Conn *net.TCPConn
	Rom  Room
}


//向用户表中添加新用户。返回true表示添加成功
func (clis *ClientsInfo) AddClient(s string, client Client) error {
	//首先判断添加的用户是否存在。如果已经存在，则返回false
	
//	if clis.clients[s] == "" {
//		return false
//	}
	clis.Clients[s] = client
	
	return nil
}

//删除用户表中的用户。返回true表示删除成功
func (clis *ClientsInfo) Deleteclient(s string) bool {

	delete(clis.Clients, s) 
	return true
}

//更新用户表中的用户。
func (clis *ClientsInfo) Updateclient(){
	
}

func NewClient(name string, conn *net.TCPConn, room Room) Client{
	addr := conn.RemoteAddr().String()
	s := strings.Split(addr, ":")
	client := Client{name, s[0], s[1], conn, room}
	return client
}

func (c *Client) SendMsg(msg []byte) {
	//room.Stat.PuSpeaker = c.Name	
	c.Rom.Stat.PuSpeaker = c.Name
	fmt.Println(c.Name)
//	sMsg = append(sMsg, []byte(c.Name)...)
//	
//	sMsg = append(sMsg, )
//	sMsg = append(sMsg, msg...)
	sMsg := []byte(c.Name + ":" + string(msg))
	c.Rom.Ch <- sMsg 
}