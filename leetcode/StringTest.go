package main

import (
	"strings"
)

func main()  {
	var str="adfaw4qfgsfdav"
 	print(lengthOfLongestSubstring(str))

}

//最长非重复子串
func lengthOfLongestSubstring(s string) int {
	if(len(s)==0) {
		return 0
	}

	start,maxLen:=0,0

	for index := 0; index < len(s); index++ {

		println(string(s[start:index]))
		println(string(s[index]))

		posIndex:=strings.Index(string(s[start:index]),string(s[index]))

		println(posIndex)

		if posIndex==-1 {
			if (maxLen<index-start+1) {
				maxLen=index-start+1
			}
		} else {
			start=start+posIndex+1
		}
	}

	return maxLen
}