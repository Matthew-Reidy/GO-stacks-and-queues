package main

import "log"

type stack []interface{}

func (s *stack) Push(element interface{}) {
	*s = append(*s, element)
}

func (s *stack) Pop() {
	if *s == nil {
		log.Panicln("ERROR: - Can not remove elements in a nil stack")
	}

	*s = (*s)[:len(*s)-1]
}

func (s *stack) Peek() interface{} {
	if *s == nil {
		log.Panicln("ERROR: - Stack is nil")
	}

	return (*s)[len(*s)-1]

}

func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}
