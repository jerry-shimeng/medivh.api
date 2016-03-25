package rwconnection


import (
	"net"
	"time"
	"rwapi/common"
	"rwapi/models"
	"fmt"
	"github.com/astaxie/beego"
)

var ConnPool map[string]*ConnectionModel


func init() {
	ConnPool = make(map[string]*ConnectionModel, 1000)
}

func AddConn(conn *net.TCPConn) {
	if conn == nil {
		return
	}
	addr := conn.RemoteAddr().String()
	ct := time.Now().Unix()
	id := common.NewObjectId().Hex()

	model := ConnectionModel{ConnAddr:addr, ConnTime:ct, ConnObj:conn, ConnId:id}
	//model.ReturnMessage = make(chan string, 10)
	ConnPool[id] = &model
	go model.BeginAccept()
	beego.BeeLogger.Info(fmt.Sprintf("添加链接：%s",conn.RemoteAddr().String()))
}


func RemoveConn(conn *net.TCPConn) {
	for i, v := range ConnPool {
		if v.ConnObj == conn {
			delete(ConnPool, i)
			beego.BeeLogger.Info("断开的链接：%s , 现有连接数量: %d",conn.RemoteAddr().String(),len(ConnPool))
		}
	}
}

func GetAllAddr() *[]models.ConnViewModel {
	var count = len(ConnPool)
	if count == 0 {
		return nil
	}
	list := make([]models.ConnViewModel, 0)

	for _, value := range ConnPool {
		if value != nil {
			temp := models.ConnViewModel{}
			temp.ConnAddr = value.ConnAddr
			temp.ConnTime = value.ConnTime
			temp.ConnId = value.AppKey
			temp.Alias = value.Alias
			list = append(list, temp)
		}
	}
	//	fmt.Printf("\n[list] %#v\n",list)

	return &list
}