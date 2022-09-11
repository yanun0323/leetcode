package main

import (
	"fmt"
)

func main() {
	// var inte interface{}
	// var boolean bool

	// fmt.Println("interface: ", unsafe.Sizeof(inte))
	// fmt.Println("bool: ", unsafe.Sizeof(boolean))
	// fmt.Println(pow(26))
	// fmt.Println(pow2(32))
	fmt.Println('a')
	fmt.Println('z')
	fmt.Println('A')
	fmt.Println('Z')
	str := "123456"
	fmt.Println(str[1:])

	a := 1
	b := 2
	c := 3

	a, b, c = c, a, b
	// 3, 3, 3
	// 3, 1, 2
	fmt.Println(a, b, c)

}
