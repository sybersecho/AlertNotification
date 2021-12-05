package model

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	Data []time.Time
	lock sync.Mutex
}

func (q *Queue) Enqueue(data time.Time) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Data = append(q.Data, data)
}

func (q *Queue) Dequeue() error {
	if len(q.Data) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()

		q.Data = q.Data[1:]
		return nil
	}

	return fmt.Errorf("pop up: queue is empty")
}

func (q *Queue) Front() (*time.Time, error) {
	if len(q.Data) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()

		return &q.Data[0], nil
	}

	return nil, fmt.Errorf("peep error: queue is empyt")
}

func (q *Queue) Last() (*time.Time, error) {
	if len(q.Data) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()

		return &q.Data[len(q.Data)-1], nil
	}

	return nil, fmt.Errorf("peep error: queue is empyt")
}

func (q *Queue) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Data = make([]time.Time, 0)
}

func (q *Queue) Size() int {
	return len(q.Data)
}

func (q *Queue) Empty() bool {
	return len(q.Data) == 0
}
