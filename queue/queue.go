package queue

type Queue struct {
	values []interface{}
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.values = append(queue.values, value)
}

func New() Queue {
	return Queue{}
}
