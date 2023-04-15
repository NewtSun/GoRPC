/**
 * @Author : NewtSun
 * @Date : 2023/4/15 12:56
 * @Description :
 **/

package client

func (call *Call) done() {
	call.Done <- call
}
