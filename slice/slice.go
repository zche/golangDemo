package main

import "fmt"

func main() {
	// var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	// var mySlice []int =myArray[:5]
	// fmt.Println("Elements of myArray:")
	// for _,v:= range myArray{
	// 	fmt.Println(v," ")
	// }
	// fmt.Println("Elements of mySlice:")
	// for _,v:= range mySlice{
	// 	fmt.Println(v," ")
	// }
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1)
	fmt.Println("Elements of slice1:")
	for _, v := range slice1 {
		fmt.Println(v, " ")
	}

	fmt.Println("Elements of slice2:")
	for _, v := range slice2 {
		fmt.Println(v, " ")
	}
	copy(slice2, slice1)
	fmt.Println("Elements of slice1:")
	for _, v := range slice1 {
		fmt.Println(v, " ")
	}

	fmt.Println("Elements of slice2:")
	for _, v := range slice2 {
		fmt.Println(v, " ")
	}

	fmt.Println()
}
