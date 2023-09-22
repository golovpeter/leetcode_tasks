package Design

import "fmt"

type MyQueue struct {
	leftStack, rightStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		leftStack:  make([]int, 0),
		rightStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.rightStack = append(this.rightStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.leftStack) > 0 {
		el := this.leftStack[len(this.leftStack)-1]
		this.leftStack = this.leftStack[:len(this.leftStack)-1]
		return el
	}

	this.leftStack = reverse(this.rightStack)
	fmt.Println(this.leftStack)
	this.rightStack = make([]int, 0)

	el := this.leftStack[len(this.leftStack)-1]
	this.leftStack = this.leftStack[:len(this.leftStack)-1]
	return el
}

func (this *MyQueue) Peek() int {
	if len(this.leftStack) > 0 {
		return this.leftStack[len(this.leftStack)-1]
	}

	this.leftStack = reverse(this.rightStack)
	this.rightStack = make([]int, 0)

	return this.leftStack[len(this.leftStack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.leftStack) == 0 && len(this.rightStack) == 0 {
		return true
	}

	return false
}

func reverse(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}

	return nums
}
