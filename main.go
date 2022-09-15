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
	fmt.Println(1 << 8)
	fmt.Println(validUtf8([]int{197, 130, 1}))
	fmt.Println(255 >> 1)
}

func validUtf8(data []int) bool {
	skip := 0
	for _, num := range data {
		v := byte(num)
		fmt.Println("v: ", v, "=========================")
		v = ^v
		fmt.Println("^v: ", v)
		if skip > 0 {
			fmt.Println("skip ^v >> 6: ", v>>6)
			if v>>6 != 1 {
				return false
			}
			skip--
			continue
		}
		v = v >> 3
		fmt.Println("^v >> 3: ", v)
		// 4, 3, 2
		for i := 0; i < 3; i++ {
			if v == 1 {
				skip = 3 - i
				fmt.Println("skip: ", skip)
				break
			}
			v = v >> 1
			fmt.Println("^v >> 1: ", v)
		}
		// 1
		if skip == 0 && v>>1 != 1 {
			return false
		}
	}
	return true
}
