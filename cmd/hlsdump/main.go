package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("./files/720.00143.ts")
	if err != nil {
		fmt.Println(err)
		return
	}
	if bytes[0] == 0x47 {
		fmt.Println("aaa")
	}
	//fmt.Println(hex.Dump(bytes))
}
