package rwconnection


import (
	"net"
	"time"
	"rwapi/common"
"rwapi/models"
)

var ConnPool map[string]*ConnectionModel


func init()  {
	ConnPool = make(map[string]*ConnectionModel,1000)
}

func AddConn(conn *net.TCPConn){
	if conn == nil {
		return
	}
	addr := conn.RemoteAddr().String()
	ct := time.Now().Unix()
	id := common.NewObjectId().Hex()

	model := ConnectionModel{ConnAddr:addr,ConnTime:ct,ConnObj:conn,ConnId:id}
	model.ReturnMessage = make(chan string,10)
	ConnPool[id] = &model
	go model.BeginAccept()

}


func RemoveConn(conn *net.TCPConn){
	for _,v := range ConnPool{
		if v.ConnObj == conn {
			delete(ConnPool,v.ConnId)
		}
	}
}

func GetAllAddr()*[]models.ConnViewModel{
	var count = len(ConnPool)
	if count == 0 {
		return nil
	}
	list := make([]models.ConnViewModel,0)

	for _,value:=range ConnPool{
		temp := models.ConnViewModel{}
		temp.ConnAddr = value.ConnAddr
		temp.ConnTime = value.ConnTime
		temp.ConnId = value.ConnId
		temp.Alias = value.Alias
		list = append(list,temp)
	}

	return &list
}

func FindConn(id string)*ConnectionModel{
	t,ok := ConnPool[id]

	if ok == false {
		return nil
	}else {
		return t
	}
}