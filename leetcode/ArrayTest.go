package main

import (
	"fmt"
	"time"
)

func main() {
	// var arr=[]int{2,3,6,4,7}
	// var target=9

	start_time:=time.Now()


	fmt.Println(logestCommonPrefix([]string{"adad","adasdas","aqq","asqwwq"}))
	

	fmt.Println(time.Since(start_time))
}

// func twoSum(nums []int,target int) []int {
// 	result := make([]int,0)
// 	for i:=0;i<len(nums);i++ {		
// 		for j:=i+1;j<len(nums);j++{
// 			if target-nums[i]==nums[j]{
// 				result=[]int{i,j}
// 				goto stop
// 			}
// 		}
// 	}

// 	stop:
// 	return result
// }

// func twoSum(nums []int, target int) []int {

// 	maps := make(map[int]int)
// 	nes := make([]int,0)
	
// 	for key,val := range nums {
// 		ant := target - val
// 		_,st := maps[ant]
// 		if st {
// 			nes = append(nes,maps[ant])
// 			nes = append(nes,key)
// 			return  nes
// 		}
// 		maps[val] = key
// 	}
// 	return nes
// }


//最长公共前缀
func logestCommonPrefix(strs []string) string {
	strsLength:=len(strs)

	if (strsLength==0) {
		return ""
	}

	maxPreLength:=len(strs[0])

	for index := 1; index < strsLength; index++ {
	strLength:=len(strs[index])

		if(strLength<maxPreLength) {
			if(strLength==0) {
				return ""
			}
			maxPreLength=strLength
		}
	}

	i:=0

	I:

	for ; i < maxPreLength; i++ {
		for j := 1; j < strsLength; j++ {
			if(strs[j][i]!=strs[j-1][i]) {
				break I
			}
		}
	}

	return strs[0][0:i]
}