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
	fmt.Println("====listener server is running====")
	start()
}


func start() {
 	var tcpAddr *net.TCPAddr
	serverAddr,_ := beego.AppConfig.Int("port")

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%d",serverAddr))
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


