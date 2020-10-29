package goqueue

import "testing"

func TestQueue(t *testing.T) {
	t.Run("it enqueues and dequeues", func(t *testing.T) {
		var queue Queue

		item1 := "thing"
		item2 := "another thing"

		queue.Enqueue(item1)
		queue.Enqueue(item2)

		item := queue.Dequeue()

		if item != item1 {
			t.Errorf("got %s want %s", item, item1)
		}

		item = queue.Dequeue()

		if item != item2 {
			t.Errorf("got %s want %s", item, item2)
		}
	})

	t.Run("it returns nil for empty dequeue", func(t *testing.T) {
		var queue Queue

		item := queue.Dequeue()

		if item != nil {
			t.Errorf("got %v want %v", item, nil)
		}
	})
}
