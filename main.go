package main

import "fmt"

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

	s := "ABCDEFG"
	fmt.Println(s[1:2])

}

func pow(time int) uint64 {
	var num uint64 = 1
	for i := 0; i < time; i++ {
		num *= 100
	}
	return num
}

func pow2(time int) int {
	num := 1
	for i := 0; i < time; i++ {
		num *= 2
	}
	return num
}
