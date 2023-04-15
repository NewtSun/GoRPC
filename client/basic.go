/**
 * @Author : NewtSun
 * @Date : 2023/4/15 13:05
 * @Description :
 **/

package client

import (
	"GoRPC/codec"
	"errors"
)

func parseOptions(opts ...*codec.Option) (*codec.Option, error) {
	// if opts is nil or pass nil as parameter
	if len(opts) == 0 || opts[0] == nil {
		return codec.DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = codec.DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = codec.DefaultOption.CodecType
	}
	return opt, nil
}
