package main

// FIXME: [0732]
type MyCalendarThree struct {
	Bookings     [][2]int
	Intersection [][2]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	current := [2]int{start, end - 1}
	currentIntersection := [2]int{1, 1}
	result := 0

	for i, booking := range this.Bookings {
		for j := 0; j <= 1; j++ {
			if isIntersect(booking[j], current[0], current[1]) {
				this.Intersection[i][j]++
			}
			result = max(result, this.Intersection[i][j])

			if isIntersect(current[j], booking[0], booking[1]) {
				currentIntersection[j]++
			}
		}
	}

	this.Bookings = append(this.Bookings, current)
	this.Intersection = append(this.Intersection, currentIntersection)

	return max(result, max(currentIntersection[0], currentIntersection[1]))
}

func isIntersect(a, b1, b2 int) bool {
	return a >= b1 && a <= b2
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
