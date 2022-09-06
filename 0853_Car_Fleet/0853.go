package main

import (
	"sort"
)

type Car struct {
	pos  int
	time float64
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, 0, len(position))
	for i := range position {
		cars = append(cars, Car{
			pos:  position[i],
			time: float64(target-position[i]) / float64(speed[i]),
		})
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})
	var tempTime float64 = 0
	result := 0
	for i := range cars {
		if cars[i].time > tempTime {
			result++
			tempTime = cars[i].time
		}
	}
	return result
}
