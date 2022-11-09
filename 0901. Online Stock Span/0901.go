package main

import "math"

type Price struct {
	val int
	i   int
}

type StockSpanner struct {
	queue []Price
	count int
}

func Constructor() StockSpanner {
	return StockSpanner{
		queue: []Price{{val: math.MaxInt, i: -1}},
		count: 0,
	}
}

func (s *StockSpanner) Next(price int) int {
	defer func() {
		s.queue = append(s.queue, Price{
			val: price,
			i:   s.count,
		})
		s.count++
	}()

	for s.queue[len(s.queue)-1].val <= price {
		s.queue = s.queue[:len(s.queue)-1]
	}

	return s.count - s.queue[len(s.queue)-1].i
}
