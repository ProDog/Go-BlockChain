package main

import (
	"fmt"
)

type MinStack struct {
	s []int
}

func mian() {

}

func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    this.s=append(this.s,x)
}


func (this *MinStack) Pop()  {
    this.s=this.s[:len(this.s)-1]
}


func (this *MinStack) Top() int {
    return this.s[len(this.s)-1]
}


func (this *MinStack) GetMin() int {
	minValue:=this.s[0]
	for i := 1; i < len(this.s); i++ {
		if(this.s[i]<minValue) {
			minValue=this.s[i]
		}
	}
	return minValue
}

