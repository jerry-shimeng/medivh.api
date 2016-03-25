package rwconnection

import (
	"net"
	"bufio"
	"fmt"
	"strings"
	"rwapi/models"
	"encoding/json"
	"errors"
	"rwapi/code/rwcache"
)

type ConnectionModel struct {
	ConnObj                  *net.TCPConn
	ConnTime                 int64
	ConnAddr                 string
	ConnId                   string
	//ReturnMessage            chan string
	Alias, AppKey, AppSecret string
}

func (this *ConnectionModel)BeginAccept() {
	reader := bufio.NewReader(this.ConnObj)
	writer := bufio.NewWriter(this.ConnObj)
	var count = 0

	defer func() {
		RemoveConn(this.ConnObj)
		writer.Write([]byte("end\n"))
		writer.Flush()
		this.ConnObj.Close()
		this.ConnObj = nil
		//close(this.ReturnMessage)
		RemoveConnDict(this.AppKey, this.ConnId)
	}()

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		//收到消息后，将消息发送到异步通道处理
		//fmt.Println(msg)
		if msg != "o\n" &&  msg != "0\n" {
			//第一次设置别名
			if count == 0 {
				count = 1
				err = this.setClientInfo(strings.Trim(msg, "\n"))
				if err != nil {
					fmt.Println("[ConnectionModel.BeginAccept]", err.Error(), msg)
					return
				}
			}else {
				//处理并反馈消息
				//添加到缓存
				rwcache.AddMsg(msg)
				//this.ReturnMessage <- msg
			}
		}
		//消息反馈
		writer.Write([]byte("o\n"))
		writer.Flush()
	}
}

func (this *ConnectionModel)setClientInfo(msg string) error {
	if len(msg) == 0 {
		return errors.New("msg is null")
	}
	var model = new(models.ClientInfo)
	err := json.Unmarshal([]byte(msg), model)
	if err != nil {
		return err
	}
	this.Alias = model.AppName
	this.AppKey = model.AppKey
	this.AppSecret = model.AppSecret
	//验证key可用性以及是否重复 和 AppSecret是否正确
	//to do list

	SetConnDict(this.ConnId, this.AppKey);
	return nil
}