package stacks_queue

//stacks ---> LIFO last in first out
//Queue ---> FIFo first in first out

type Stack struct {
	array []any
}

//Push
func (s *Stack) Push(val any) {
	s.array = append(s.array, val)
}

//Pop
func (s *Stack) Pop() {
	if len(s.array) == 0 {
		return
	}
	s.array = s.array[:len(s.array)-1]
}

type Queue struct {
	array []any
}

func (q *Queue) Enqueue(val any) {
	q.array = append(q.array, val)
}

func (q *Queue) Dequeue() {
	if len(q.array) == 0 {
		return
	}
	q.array = q.array[1:]
}
