/**
 * @Author : NewtSun
 * @Date : 2023/4/15 11:20
 * @Description :
 **/

package main

import (
	"GoRPC/client"
	"GoRPC/server"
	"time"
)

func main() {
	//server.StartServer()
	go server.StartServer(":8888")
	go server.StartServer(":9999")

	time.Sleep(3 * time.Second)

	//client.StartClient()

	client.StartXClient(":8888", ":9999")

}
