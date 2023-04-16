/**
 * @Author : NewtSun
 * @Date : 2023/4/15 11:12
 * @Description :
 **/

package server

import (
	"GoRPC/codec"
	"log"
	"net"
)

//func StartServer(addr ...chan string) {
//	// pick a free port
//	l, err := net.Listen("tcp", ":8888")
//	if err != nil {
//		log.Fatal("network error:", err)
//	}
//	log.Println("start rpc server on", l.Addr())
//	//addr <- l.Addr().String()
//	Accept(l)
//}

func StartServer(addr string, addrCh ...chan string) {
	var foo codec.Foo
	//if err := Register(&foo); err != nil {
	//	log.Fatal("register error:", err)
	//}
	_ = Register(&foo)
	// pick a free port ":8888"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	//addr <- l.Addr().String()
	Accept(l)
}
