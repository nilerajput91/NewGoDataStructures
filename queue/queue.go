package main

import "fmt"

// Queue represents a queue that holds the slice
type Queue struct {
	items []int
}

// Enqueue adds the value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dqueue remove the value form queue
func (q *Queue) Dqueue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(30)

	fmt.Println(myQueue)

	myQueue.Dqueue()
	fmt.Println(myQueue)
	myQueue.Dqueue()
	fmt.Println(myQueue)
}
