/*
服务监听
*/
package rwnetwork

import (
	"fmt"
	"net"
	"rwapi/code/rwconnection"
	"github.com/astaxie/beego"
)

func Run() {
	fmt.Println("listener server is running")
	start()
}

//var connObj *net.TCPConn
//var receiveMsg chan string


func start() {
 	var tcpAddr *net.TCPAddr
	serverAddr := beego.AppConfig.String("host")
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("[ERR]", err)
	}
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err.Error())
		}
		rwconnection.AddConn(tcpConn)

	}
}

//func Accept(tcpConn *net.TCPConn) {
//	reader := bufio.NewReader(tcpConn)
//	writer := bufio.NewWriter(tcpConn)
//
//	for {
//		msg, err := reader.ReadString('\n')
//		if err != nil {
//			fmt.Println(err)
//			rwconnection.RemoveConn(tcpConn)
//			tcpConn.Close()
//			break
//		}
//		//收到消息后，将消息发送到异步通道处理
//		//fmt.Println(msg)
//		if msg != "o\n" {
//			//处理并反馈消息
//			receiveMsg <- msg
//		}
//		//消息反馈
//		writer.Write([]byte("o\n"))
//		writer.Flush()
//	}
//}
//
////发送命令
//func SendCmd(s *string,tcpConn *net.TCPConn)error{
//	if s == nil {
//		return errors.New("cmd is nil")
//	}
//	if connObj == nil {
//		return errors.New("connection is error")
//	}
//	writer := bufio.NewWriter(tcpConn)
//	writer.Write([]byte(*s))
//	return writer.Flush()
//}
//
//func Receive()*string{
//
//	t1 := time.After(10*time.Second)
//
//	select {
//		case s :=  <-receiveMsg:
//			return &s
//		case <-t1:
//			return nil
//	}
//}

