package main

import "fmt"

// simple queue implement by go

type Queue struct {
	items []string
}

func (q *Queue) EnQueue(data string) {
	q.items = append(q.items, data)
}

func (q *Queue) DeQueue() string {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return "null"
	}
	data := q.items[0]
	q.items = q.items[1:]
	return data

}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
