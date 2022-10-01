package main

type MyCircularQueue struct {
	queue []int
	cap   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, 0, k),
		cap:   k,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.queue = append(q.queue, value)
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.queue = q.queue[1:]
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[0]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[len(q.queue)-1]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return len(q.queue) == q.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
