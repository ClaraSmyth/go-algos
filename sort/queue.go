package sort

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head, tail *Node[T]
	length     int
}

func (q *Queue[T]) enqueue(val T) {
	q.length += 1
	if q.tail == nil {
		q.head = &Node[T]{value: val}
		q.tail = q.head
	} else {
		q.tail.next = &Node[T]{value: val}
		q.tail = q.tail.next
	}
}

func (q *Queue[T]) dequeue() (T, bool) {
	if q.head == nil {
		var v T
		return v, false
	}

	currentHead := q.head
	q.length -= 1
	q.head = q.head.next

	return currentHead.value, true
}

func (q *Queue[T]) peek() T {
	return q.head.value
}

func TestQueue() string {
	q := Queue[int]{}

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)

	test1 := q.peek()
	test2 := q.length
	test3, _ := q.dequeue()
	test4 := q.length

	q.dequeue()
	q.dequeue()
	_, exists := q.dequeue()

	if test1 == 1 && test2 == 3 && test3 == 1 && test4 == 2 && !exists {
		return "Pass"
	}
	return "Fail"
}
