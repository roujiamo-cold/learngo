package main

import (
	"fmt"
	"learngo/imooc/google/lang/container/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//q.Push("asd")
	//fmt.Println(q.Pop())
}
