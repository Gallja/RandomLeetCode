package main

import (
	"fmt"
)

type LinkedList struct {
	head *ListNode
	size int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *LinkedList) add(Val int) {
	newNode := &ListNode{Val, nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.Next != nil {
			current = current.Next
		}

		current.Next = newNode
	}

	list.size++
}

func (list *LinkedList) printList() {
	if list.size == 0 {
		fmt.Println("[ ]")
		return
	}

	fmt.Print("[ ", list.head.Val)

	current := list.head.Next

	for current != nil {
		fmt.Print(", ", current.Val)
		current = current.Next
	}

	fmt.Print(" ]")
	fmt.Println()
}

func doubleIt(head *ListNode) *ListNode {
	if head.Val == 0 && head.Next == nil {
		listRet := LinkedList{nil, 0}
		listRet.add(0)

		return listRet.head
	}

	nums := []int{}
	var num int

	current := head
	cnt := 0

	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
		cnt++
	}

	fmt.Println(nums)

	for _, Val := range nums {
		num = num*10 + Val
	}

	fmt.Println(num)

	num *= 2

	fmt.Println(num)

	slice := intToSlice(num)

	listRet := LinkedList{nil, 0}

	for i := 0; i < len(slice); i++ {
		listRet.add(slice[i])
	}

	return listRet.head
}

func intToSlice(n int) []int {
	var result []int

	for n > 0 {
		digit := n % 10
		result = append([]int{digit}, result...)
		n /= 10
	}

	return result
}

func main() {
	list := LinkedList{nil, 0}

	list.add(9)
	list.add(1)
	list.add(9)
	list.add(5)
	/*
		list.add(0)
		list.add(5)
		list.add(1)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
		list.add(9)
	*/

	list.printList()

	fmt.Println(doubleIt(list.head))
}
