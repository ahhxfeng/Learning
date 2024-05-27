package main

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	// test Dequeue when queue is empty
	t.Log("Test queue start ")
	Sq := Queue{}
	Sq.DeQueue()

	Sq.EnQueue("今天")
	Sq.EnQueue("天气")
	Sq.EnQueue("怎么样")
	Sq.EnQueue("?")

	for !Sq.IsEmpty() {
		fmt.Println(Sq.DeQueue())
	}
}
