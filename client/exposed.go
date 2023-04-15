/**
 * @Author : NewtSun
 * @Date : 2023/4/15 13:14
 * @Description :
 **/

package client

import (
	"GoRPC/codec"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

//// Dial connects to an RPC server at the specified network address
//func Dial(network, address string, opts ...*codec.Option) (client *Client, err error) {
//	opt, err := parseOptions(opts...)
//	if err != nil {
//		return nil, err
//	}
//	conn, err := net.Dial(network, address)
//	if err != nil {
//		return nil, err
//	}
//	// close the connection if client is nil
//	defer func() {
//		if client == nil {
//			_ = conn.Close()
//		}
//	}()
//	return NewClient(conn, opt)
//}

func dialTimeout(f newClientFunc, network, address string, opts ...*codec.Option) (client *Client, err error) {
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTimeout(network, address, opt.ConnectTimeout)
	if err != nil {
		return nil, err
	}
	// close the connection if client is nil
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()
	ch := make(chan clientResult)
	go func() {
		client, err := f(conn, opt)
		ch <- clientResult{client: client, err: err}
	}()
	if opt.ConnectTimeout == 0 {
		result := <-ch
		return result.client, result.err
	}
	select {
	case <-time.After(opt.ConnectTimeout):
		return nil, fmt.Errorf("rpc client: connect timeout: expect within %s", opt.ConnectTimeout)
	case result := <-ch:
		return result.client, result.err
	}
}

// Dial connects to an RPC server at the specified network address
func Dial(network, address string, opts ...*codec.Option) (*Client, error) {
	return dialTimeout(NewClient, network, address, opts...)
}

// Go invokes the function asynchronously.
// It returns the Call structure representing the invocation.
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	client.send(call)
	return call
}

//// Call invokes the named function, waits for it to complete,
//// and returns its error status.
//func (client *Client) Call(serviceMethod string, args, reply interface{}) error {
//	call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
//	return call.Error
//}

// Call invokes the named function, waits for it to complete,
// and returns its error status.
func (client *Client) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	call := client.Go(serviceMethod, args, reply, make(chan *Call, 1))
	select {
	case <-ctx.Done():
		client.removeCall(call.Seq)
		return errors.New("rpc client: call failed: " + ctx.Err().Error())
	case call := <-call.Done:
		return call.Error
	}
}
