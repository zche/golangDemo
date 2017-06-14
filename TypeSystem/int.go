package main

import "fmt"

//Integer 表示一个整形。。
type Integer int

//Less 比较大小
func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

}
