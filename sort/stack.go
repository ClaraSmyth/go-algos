package sort

type SNode[T any] struct {
	value T
	prev  *SNode[T]
}

type Stack[T any] struct {
	head   *SNode[T]
	length int
}

func (s *Stack[T]) push(value T) {
	s.length += 1
	if s.head == nil {
		s.head = &SNode[T]{value: value}
	} else {
		currentHead := s.head
		s.head = &SNode[T]{value: value}
		s.head.prev = currentHead
	}
}

func (s *Stack[T]) pop() (T, bool) {
	if s.head == nil {
		var v T
		return v, false
	}

	s.length -= 1
	currentHead := s.head
	s.head = s.head.prev
	return currentHead.value, true
}

func (s *Stack[T]) peek() (T, bool) {
	if s.head == nil {
		var v T
		return v, false
	}

	return s.head.value, true
}

func TestStack() string {
	s := &Stack[int]{}

	s.push(1)
	s.push(2)
	s.push(3)

	test1, _ := s.peek()
	test2 := s.length
	test3, _ := s.pop()
	test4 := s.length
	test5, _ := s.peek()

	s.pop()
	s.pop()
	_, exists := s.pop()

	if test1 == 3 && test2 == 3 && test3 == 3 && test4 == 2 && test5 == 2 && !exists {
		return "Pass"
	}

	return "Fail"
}
