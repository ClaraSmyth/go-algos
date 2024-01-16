package list

import (
	"errors"
)

type LNode[T comparable] struct {
	value T
	next  *LNode[T]
	prev  *LNode[T]
}

type LinkedList[T comparable] struct {
	head   *LNode[T]
	tail   *LNode[T]
	length int
}

func (list *LinkedList[T]) prepend(val T) {
	newNode := &LNode[T]{value: val}
	list.length++

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
}

func (list *LinkedList[T]) append(val T) {
	newNode := &LNode[T]{value: val}
	list.length++

	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *LinkedList[T]) insertAt(val T, idx int) error {

	if idx > list.length {
		return errors.New("index out of range")
	} else if idx == 0 {
		list.prepend(val)
	} else if idx == list.length {
		list.append(val)
	} else {
		newNode := &LNode[T]{value: val}
		list.length++

		currNode := list.head

		for i := 0; i < idx; i++ {
			currNode = currNode.next
		}

		newNode.next = currNode
		newNode.prev = currNode.prev
		currNode.prev.next = newNode
		currNode.prev = newNode
	}

	return nil
}

func (list *LinkedList[T]) remove(val T) error {
	if list.length == 0 {
		return errors.New("list is empty")
	}

	currNode := list.head

	for i := 0; i < list.length; i++ {
		if currNode.value == val {
			break
		} else if currNode.next == nil {
			return errors.New("value doesnt exist")
		}
		currNode = currNode.next
	}

	list.length--

	if list.length == 0 {
		list.head = nil
		list.tail = nil
	} else if currNode.prev == nil {
		list.head = currNode.next
	} else if currNode.next == nil {
		list.tail = currNode.prev
	} else {
		currNode.prev.next = currNode.next
		currNode.next.prev = currNode.prev
	}

	return nil
}

func (list *LinkedList[T]) valueAt(idx int) (T, error) {
	if list.length == 0 || idx > list.length-1 {
		var v T
		return v, errors.New("out of range")
	}

	currNode := list.head

	for i := 0; i < idx; i++ {
		currNode = currNode.next
	}

	return currNode.value, nil
}

func TestLinkedList() string {
	list := &LinkedList[int]{}

	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.prepend(5)

	test1 := list.length
	test2, _ := list.valueAt(0)

	list.insertAt(6, 2)

	test3, _ := list.valueAt(2)
	test4, _ := list.valueAt(3)

	list.remove(5)

	test5 := list.length
	test6, _ := list.valueAt(0)

	if test1 == 5 && test2 == 5 && test3 == 6 && test4 == 2 && test5 == 5 && test6 == 1 {
		return "Pass"
	}

	return "Fail"
}
