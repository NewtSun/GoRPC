/**
 * @Author : NewtSun
 * @Date : 2023/4/15 11:09
 * @Description :
 **/

package codec

type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // sequence number chosen by client
	Error         string
}
