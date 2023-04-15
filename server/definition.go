/**
 * @Author : NewtSun
 * @Date : 2023/4/15 11:17
 * @Description :
 **/

package server

import (
	"GoRPC/codec"
	"reflect"
	"sync"
)

// Server represents an RPC Server.
type Server struct {
	serviceMap sync.Map
}

type request struct {
	h            *codec.Header // header of request
	argv, replyv reflect.Value // argv and replyv of request
	mtype        *methodType
	svc          *service
}

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
	numCalls  uint64
}

type service struct {
	name    string
	typ     reflect.Type
	rcvr    reflect.Value
	methods map[string]*methodType
}
