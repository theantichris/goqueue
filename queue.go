package goqueue

import (
	"fmt"
	"sync"
)

// Item represents a single item in the queue.
type Item interface{}

// Queue is the queue data structure that holds the items.
type Queue struct {
	items []Item
	mutex sync.Mutex
}

// Enqueue adds an item to the rear of the queue.
func (q *Queue) Enqueue(item Item) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
func (q *Queue) Dequeue() Item {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {
		fmt.Println("the queue is empty")

		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item
}
