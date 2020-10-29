package goqueue

import "testing"

const (
	item1 = "thing"
	item2 = "another thing"
)

func TestEnqueue(t *testing.T) {
	t.Run("it enqueues items", func(t *testing.T) {
		var queue Queue

		queue.Enqueue(item1)
		queue.Enqueue(item2)

		if queue.items[0] != item1 {
			t.Errorf("got %s want %s", queue.items[0], item1)
		}

		if queue.items[1] != item2 {
			t.Errorf("got %s want %s", queue.items[1], item2)
		}
	})
}

func TestDequeue(t *testing.T) {
	t.Run("it dequeues items", func(t *testing.T) {
		var queue Queue

		queue.Enqueue(item1)
		queue.Enqueue(item2)

		got := queue.Dequeue()

		if got != item1 {
			t.Errorf("got %v want %v", got, item1)
		}

		got = queue.Dequeue()

		if got != item2 {
			t.Errorf("got %v want %v", got, item2)
		}
	})

	t.Run("it returns nil for empty queue", func(t *testing.T) {
		var queue Queue

		item := queue.Dequeue()

		if item != nil {
			t.Errorf("got %v want %v", item, nil)
		}
	})
}

func TestReset(t *testing.T) {
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
}

func TestDump(t *testing.T) {
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

func TestPeek(t *testing.T) {
	t.Run("it returns the first item", func(t *testing.T) {
		t.Run("it returns the first item in the queue", func(t *testing.T) {
			var queue Queue

			queue.Enqueue(item1)
			queue.Enqueue(item2)

			if queue.Peek() != item1 {
				t.Errorf("got %v want %v", queue.Peek(), item1)
			}

			if queue.Peek() != item1 {
				t.Errorf("got %v want %v", queue.Peek(), item1)
			}
		})
	})

	t.Run("it returns nil if the queue is empty", func(t *testing.T) {
		var queue Queue

		if queue.Peek() != nil {
			t.Errorf("got %v want %v", queue.Peek(), nil)
		}
	})
}
