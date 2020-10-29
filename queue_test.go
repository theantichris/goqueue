package goqueue

import "testing"

const (
	item1 = "thing"
	item2 = "another thing"
)

func TestQueue(t *testing.T) {
	t.Run("it enqueues and dequeues", func(t *testing.T) {
		var queue Queue

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

	t.Run("it resets the queue", func(t *testing.T) {
		var queue Queue

		queue.Enqueue(item1)
		queue.Enqueue(item2)

		if queue.IsEmpty() {
			t.Fatal("items did not enqueue")
		}

		queue.Reset()

		if !queue.IsEmpty() {
			t.Error("queue did not reset")
		}
	})

	t.Run("it dumps the stack", func(t *testing.T) {
		var queue Queue

		queue.Enqueue(item1)
		queue.Enqueue(item2)

		got := queue.Dump()

		want := []Item{}
		want = append(want, item1)
		want = append(want, item2)

		if got[0] != want[0] {
			t.Errorf("got %v want %v", got[0], want[0])
		}

		if got[1] != want[1] {
			t.Errorf("got %v want %v", got[1], want[1])
		}
	})
}
