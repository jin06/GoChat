package chat

type Room struct {
	Num    int
	Name   string
	Client map[string]Client
	Ch     chan []byte
	Stat   *Status
}

//房间状态信息

type Status struct {
	PuSpeaker string //当前谁在公共区域说话
}

func NewRoom(i int, name string) Room {
	return Room{i, name, make(map[string]Client), make(chan []byte), &Status{""}}
}

//添加用户到Room中
func (room Room) AddClient(c Client) {
	room.Client[c.Name] = c
	//初次进入聊天室，打印欢迎信息
	//	welcom := c.Name + "! 欢迎来到" + string(room.Num) + "聊天室"
	//	c.Conn.Write([]byte(welcom))
}

//删除room中的用户
func (room Room) RemoveClient(c Client) {
	delete(room.Client, c.Name)
}

//用户从froRoom转移到toRoom
func (froRoom Room) UpdateClent(c Client, toRoom *Room) {
	froRoom.RemoveClient(c)
	toRoom.AddClient(c)
}

//房间启动，为用户打印聊天信息等
func (room Room) Start() {

	for {

		message := <-room.Ch
		for k, v := range room.Client {
			if room.Stat.PuSpeaker != k {
				v.Conn.Write(message)
				//fmt.Println(room.Stat.PuSpeaker, k)
			}
		}
	}

}
