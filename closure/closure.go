package main

import "fmt"

func main() {
	var j = 5
	a := func() func() {
		var i = 10
		return func() {
			fmt.Printf("i,j:%d,%d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}
