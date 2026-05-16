package main

import (
	"fmt"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func SolvePartOne() int {
	arr := []int{0}
	val := 1
	i := 0
	for range 2017 {
		i += 376
		i %= len(arr)

		arr = slices.Insert(arr, i+1, val)
		i++
		val++

	}

	if i == 2017 {
		return arr[0]
	}

	return arr[i+1]
}

func SolvePartOneLinkedList() int {

	root := &ListNode{
		Val:  0,
		Next: nil,
	}
	root.Next = root

	val := 1
	curr := root
	for range 2017 {
		for range 376 {
			curr = curr.Next
		}

		nxt := &ListNode{
			Val:  val,
			Next: curr.Next,
		}

		curr.Next = nxt
		curr = curr.Next
		val++

	}

	return curr.Next.Val

}

func main() {
	res := SolvePartOne()
	res = SolvePartOneLinkedList()
	fmt.Println(res)

}
