package Stack

import (
	"math"
)

type MinStack struct {
	Stack        []int
	MinimalStack []int
	MinEl        int
}

func Constructor() MinStack {
	return MinStack{
		Stack:        make([]int, 0),
		MinimalStack: make([]int, 0),
		MinEl:        math.MaxInt64,
	}
}

func (this *MinStack) Push(val int) {
	if this.MinEl >= val {
		this.MinEl = val
		this.MinimalStack = append(this.MinimalStack, val)
	}

	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	if this.Stack[len(this.Stack)-1] == this.MinEl {

		if len(this.MinimalStack) > 1 {
			this.MinEl = this.MinimalStack[len(this.MinimalStack)-2]
		}

		this.MinimalStack = this.MinimalStack[:len(this.MinimalStack)-1]
	}

	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinEl
}
