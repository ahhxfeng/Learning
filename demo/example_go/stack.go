package main

import (
	"fmt"
)

// simple implentment about stack in go

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)

}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return "null"
	}

	data := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]
	// fmt.Printf("%s", data)
	return data
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
