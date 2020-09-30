package main

import (
	"io"
	"log"
	"net"
	"os"
)

/**
tcp的客户端
 */
func main() {
	//创建一个tcp的实例
	con,err := net.Dial("tcp","127.0.0.1:8080")

	//判断错误
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	mustCopy(os.Stdout, con)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}