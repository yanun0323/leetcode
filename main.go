package main

import "fmt"

type Room struct {
	Seats [][]int
}

func (room Room) Set(l, r, val int) {

}

func main() {
	sli := []int{0, 0, 0, 0, 0}
	sli2D := [][]int{{0, 0, 0, 0, 0}}

	func(sli []int) {
		c := make([]int, 0, len(sli))
		copy(c, sli)
		for i := range c {
			c[i] = 7
		}
	}(sli)

	func(sli2D [][]int) {
		c := make([][]int, 0, len(sli))
		copy(c, sli2D)
		for i := range c {
			for j := range c[i] {
				c[i][j] = 7
			}
		}
	}(sli2D)

	fmt.Println(sli)
	fmt.Println(sli2D)
}
