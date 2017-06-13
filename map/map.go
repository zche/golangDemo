package main

import "fmt"

//PersonInfo 是一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	//var personDB map[string]PersonInfo
	var personDB = make(map[string]PersonInfo)
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203"}
	personDB["1"] = PersonInfo{"1", "jack", "Room 101"}
	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("没找到ID为1234的人")
	}
}
