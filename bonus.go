package main

import "fmt"

type Node struct {
	value   int
	pointer *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("В стэке ничего нет")
	}
	value := s.head.value
	s.head = s.head.pointer
	s.size--
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("В стэке ничего нет")
	}
	return s.head.value, nil
}

func (s *Stack) Clear() {
	s.head = nil
	s.size = 0
}

func (s *Stack) Contains(value int) bool {
	node := s.head
	for node != nil {
		if node.value == value {
			return true
		}
		node = node.pointer
	}
	return false
}

func (s *Stack) Increment() {
	node := s.head
	for node != nil {
		node.value++
		node = node.pointer
	}
}

func (s *Stack) PrintReverse() {
	node := s.head
	var stack []int

	for node != nil {
		stack = append(stack, node.value)
		node = node.pointer
	}

	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}
}

func main() {
	s := Stack{}

	s.Push(22)
	s.Push(21)
	s.Push(212)
	s.Push(44)

	if s.Contains(10) {
		fmt.Println("В стэке есть 10")
	} else {
		fmt.Println("В стэке нет 10")
	}
	s.Increment()
	s.PrintReverse()
	s.Clear()
	fmt.Printf("Размер стэка: %d\n", s.size)

}
