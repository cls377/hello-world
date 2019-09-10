package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	Start(os.Args[1]) //接收终端信息
}
func Start(tcpAddrStr string) {
	//1.根据输入的IP加端口生成TCP的ADD信息
	//ResolveTCPAddr用于获取一个TCPAddr
	//network参数是"tcp4"、"tcp6"、"tcp"
	//tcpAddrStr表示域名或IP地址加端口号(字符串)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpAddrStr)
	if err != nil {
		log.Printf("Resolve tcp addr failed: %v\n", err)
		return
	}
	//2.向服务器拨号
	//DialTCP建立一个TCP连接
	//net参数是"tcp4"、"tcp6"、"tcp"
	//laddr表示本机地址，一般设为nil
	//raddr表示远程地址
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Dial to server failed: %v\n", err)
		return
	}
	//3:开启协程发送消息，因为输入会阻塞，所以使用for循环
	go SendMsg(conn)

	//4.接收来自服务器的广播消息
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("recv server msg failed:%v\n", err)
			conn.Close()
			os.Exit(0)
			break
		}
		fmt.Println(string(buf[0:length]))
	}
}
func SendMsg(conn net.Conn) {
	//username := conn.LocalAddr().String()
	username := "ni哥"
	for {
		var input string
		//接收输入消息，放到input变量中
		fmt.Scanln(&input)
		if input == "/q" || input == "/quit" {
			fmt.Println("Byebye...")
			conn.Close()
			os.Exit(0)
		}
		//只处理有内容端消息
		if len(input) > 0 {
			msg := "Client " + username + " say:" + input
			//向conn中写入数据
			_, err := conn.Write([]byte(msg))
			if err != nil {
				conn.Close()
				break
			}
		}
	}
}

//陈陆帅骚儿