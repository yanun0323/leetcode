package main

import (
	"fmt"
	"sync"
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
	ch := make(chan int, 40)
	wg := &sync.WaitGroup{}
	wg.Add(40)
	for i := 0; i < 40; i++ {
		ch <- i
		wg.Done()
	}
	wg.Wait()

	fmt.Println("len: ", len(ch))
	wg.Add(len(ch))
	c := len(ch)
	for i := 0; i < c; i++ {
		e := <-ch
		if e > 10 {
			ch <- e
		}
		wg.Done()
	}
	wg.Wait()

	wg.Add(len(ch))
	c = len(ch)
	fmt.Println("len: ", len(ch))
	for i := 0; i < c; i++ {
		fmt.Println(<-ch)
		wg.Done()
	}
	wg.Wait()

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
