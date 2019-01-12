package queue

type Queue struct {
	values []interface{}
}

func New() Queue {
	return Queue{}
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.values = append(queue.values, value)
}

func (queue *Queue) Dequeue() interface{} {
	newQueue := queue.values[1:]
	dequeueValue := queue.values[0]
	queue.values = nil
	queue.values = newQueue
	return dequeueValue
}
