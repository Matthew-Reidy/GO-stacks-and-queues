package main

import "log"

type queue []interface{}

func (q *queue) peek() interface{} {
	if q == nil {
		return nil
	}

	return (*q)[0]

}

func (q *queue) remove() {
	if q == nil {
		log.Panicln("ERROR: Can not remove elements in a nil Queue")
	}

	*q = (*q)[1:]
}

func (q *queue) add(element interface{}) {
	el := make(queue, 1)
	el[0] = element
	*q = append(el, *q...)

}
