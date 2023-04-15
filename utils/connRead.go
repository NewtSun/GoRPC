/**
 * @Author : NewtSun
 * @Date : 2023/4/15 13:37
 * @Description :
 **/

package utils

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const (
	CONNREAD  = "conn.Read"
	BUFIO     = "bufio.Read"
	BUFIOTEXT = "NEW_CONNECTION"
)

func ConnRead(conn io.ReadWriteCloser, rtype string) {
	switch rtype {
	case CONNREAD:
		for {
			buf := make([]byte, 1024)
			fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.(*net.TCPConn).RemoteAddr().String())
			n, err := conn.Read(buf) //从 conn 读取
			if err != nil {
				fmt.Printf("客户端退出 err=%v\n", err)
				return
			}
			fmt.Print(string(buf[:n]))
		}
	case BUFIO:
		reader := bufio.NewReader(conn)
		content, err := reader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println(string(content))

	}
}
