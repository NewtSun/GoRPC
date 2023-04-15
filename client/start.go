/**
 * @Author : NewtSun
 * @Date : 2023/4/15 13:16
 * @Description :
 **/

package client

import (
	"GoRPC/codec"
	"context"
	"log"
	"sync"
)

//func StartClient() {
//	//log.SetFlags(0)
//	client, _ := Dial("tcp", ":8888")
//	defer func() { _ = client.Close() }()
//
//	log.Println("start rpc client dial on :8888")
//	// send request & receive response
//	var wg sync.WaitGroup
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			args := fmt.Sprintf("gorpc req %d\n", i)
//			var reply string
//			if err := client.Call("Foo.Sum", args, &reply); err != nil {
//				log.Fatal("call Foo.Sum error:", err)
//			}
//			log.Println("reply:", reply)
//		}(i)
//	}
//	wg.Wait()
//}

func StartClient() {
	//log.SetFlags(0)
	client, _ := Dial("tcp", ":8888")
	defer func() { _ = client.Close() }()

	// send request & receive response
	log.Println("start rpc client dial on :8888")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &codec.Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call(context.Background(), "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}
