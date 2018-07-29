package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	result := new(ListNode)
	result.Val = 50
	tmp1 := *result
	tmp2 := result
	tmp3 := new(ListNode)
	*tmp3 = *result
	result.Val = 100
	fmt.Println("result", result.Val)
	fmt.Println("tmp1=*result", tmp1.Val)
	fmt.Println("tmp2=result", tmp2.Val)
	fmt.Println("*tmp3=*result", tmp3.Val)
}
