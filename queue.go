package main

import (
	"errors"
)

type queue []interface{}

func (q *queue) peek() (interface{}, error) {
	if q == nil {
		return nil, errors.New("Can not retrieve element of a nil queue")
	}

	return (*q)[0], nil

}

func (q *queue) remove() error {
	if q == nil {
		return errors.New("Can not remove element of a nil queue")
	}

	*q = (*q)[1:]
	return nil
}

func (q *queue) add(element interface{}) {
	el := make(queue, 1)
	el[0] = element
	*q = append(el, *q...)

}
