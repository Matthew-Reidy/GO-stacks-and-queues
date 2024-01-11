package main

type stack []interface{}

func (s *stack) Push(element interface{}) {
	*s = append(*s, element)
}

func (s *stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) Peek() interface{} {
	peeked := (*s)[len(*s)-1]
	return peeked

}

func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}
