package main

import "errors"

type stack []interface{}

func (s *stack) Push(element interface{}) {
	*s = append(*s, element)
}

func (s *stack) Pop() error {
	if *s == nil {
		return errors.New("ERROR: - Can not remove elements in a nil stack")
	}

	*s = (*s)[:len(*s)-1]
	return nil
}

func (s *stack) Peek() (interface{}, error) {
	if *s == nil {
		return nil, errors.New("ERROR: - Stack is nil")
	}

	return (*s)[len(*s)-1], nil

}

func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}
