package main

import (
	"io"
	"log"
	"net"
	"time"
)
/**
	说明：
   网络编程是并发大显身手的一个领域，由于服务器是最典型的需要同时处理很多连接的程序，这些连接一般来自于彼此独立的客户端。
 */

func main() {
	//开启一个tcp监听8080端口
	listen,err := net.Listen("tcp","127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		//获取一个连接
		con,e := listen.Accept()

		//判断是否成功
		if e != nil {
			log.Println(e)
			continue
		}

		//处理请求---开启并发
		go handel(con)
	}

}
/**
	处理
 */
func handel(c net.Conn) {
	//延时函数
	defer c.Close()

	//写如缓存
	for {
		_,err := io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
