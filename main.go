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
	//server.RunServer()
	go server.RunServer()

	time.Sleep(5 * time.Second)

	client.StartClient()

}
