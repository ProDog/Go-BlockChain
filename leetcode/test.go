package main

import (
	"fmt"
	"time"
)

func main() {
	var arr=[]int{2,3,6,4,7}
	var target=9

	start_time:=time.Now()


	fmt.Println(twoSum(arr,target))
	

	fmt.Println(time.Since(start_time))
}

func twoSum(nums []int,target int) []int {
	result := make([]int,0)
	for i:=0;i<len(nums);i++ {		
		for j:=i+1;j<len(nums);j++{
			if target-nums[i]==nums[j]{
				result=[]int{i,j}
				goto stop
			}
		}
	}

	stop:
	return result
}