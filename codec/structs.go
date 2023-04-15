/**
 * @Author : NewtSun
 * @Date : 2023/4/15 14:56
 * @Description :
 **/

package codec

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}
