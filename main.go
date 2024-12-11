package main

import (
	// doublylinkedlist "dsa-in-go/doubly-linkedlist"
	"dsa-in-go/queue"
	"fmt"
	// linkedlist "dsa-in-go/linked-list"
	// stack "dsa-in-go/stack"
)

func main() {
	fmt.Println("Queue")
	q := queue.MakeQueue[int](5)
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	q.Add(5)

	fmt.Println(q.Delete())
	// stk := stack.MakeStack[int]()
	// stk.Push(1)
	// fmt.Println(stk.Peek())
	// fmt.Println(stk.IsEmpty())
	// fmt.Println(stk.Pop())
	// fmt.Print(stk.IsEmpty())
	// fmt.Println("Doubly Linked List")
	// dll := doublylinkedlist.CreateDoublyLinkedList()
	// // dll.AddFirst(1)
	// // dll.AddFirst(2)
	// // dll.AddFirst(3)
	// // dll.AddLast(4)
	// err := dll.DeleteAtPos(0)
	// if err != nil {
	//   fmt.Print("Error :",err)
	// }
	// dll.PrintLinkedList()
	//
	// fmt.Println("LINKED LIST")
	// // head = head.AddFirst(0)
	// // head.AddAtPosition(2,12)
	// // head.AddLast(11)
	// // head.PrintLinkedList()
	// // fmt.Println("\nLength of linkedlist is ",head.GetLength())
}
