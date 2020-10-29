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

	if q.IsEmpty() {
		fmt.Println("the queue is empty")

		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item
}

// IsEmpty returns true if there are no items in the queue.
// It returns false otherwise.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Reset removes all items from the queue
func (q *Queue) Reset() {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = nil
}

// Dump returns all the items in the queue.
func (q *Queue) Dump() []Item {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	var copied = make([]Item, len(q.items))
	copy(copied, q.items)

	return copied
}

// Peek returns the first item in the queue but does not remove it.
func (q *Queue) Peek() Item {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.IsEmpty() {
		fmt.Println("the queue is empty")

		return nil
	}

	return q.items[0]
}
