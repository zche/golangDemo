package mymath

import "errors"
import "fmt"

//Add 表示两个非负数相加
func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers")
		return
	}
	return a + b, nil
}

func ManyArgsFunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
