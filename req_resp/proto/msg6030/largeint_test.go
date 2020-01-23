package msg6030

import (
	"fmt"
	"testing"
)

func TestLargeInt(t *testing.T) {
	var a uint32 = 1073741823
	var b uint32 = 1073741824
	var c uint32 = b*2
	var d uint32 = b*3 + a
	fmt.Println(a)
	fmt.Println(a & 1)
	fmt.Println(a >> 30)
	fmt.Println(b >> 30)
	fmt.Println(c >> 30)
	fmt.Printf("%032b\n",a)
	fmt.Printf("%032b\n",b)
	fmt.Printf("%032b\n",c)
	fmt.Printf("%032b\n",d)
	fmt.Printf("%064b\n",ConvertToInt(a))
	fmt.Printf("%064b\n",ConvertToInt(b))
	fmt.Printf("%064b\n",ConvertToInt(c))
	fmt.Printf("%064b\n",ConvertToInt(d))

}