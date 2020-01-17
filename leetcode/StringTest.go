package main

import (
	//"strings"
	"fmt"
	"time"
)

func main() {
	//var str="adfaw4qfgsfdav"
 	//print(lengthOfLongestSubstring(str))

	start_time:=time.Now()


	fmt.Println(mySqrt(8))
	

	fmt.Println(time.Since(start_time))

}

// //最长非重复子串
// func lengthOfLongestSubstring(s string) int {
// 	if(len(s)==0) {
// 		return 0
// 	}

// 	start,maxLen:=0,0

// 	for index := 0; index < len(s); index++ {

// 		println(string(s[start:index]))
// 		println(string(s[index]))

// 		posIndex:=strings.Index(string(s[start:index]),string(s[index]))

// 		println(posIndex)

// 		if posIndex==-1 {
// 			if (maxLen<index-start+1) {
// 				maxLen=index-start+1
// 			}
// 		} else {
// 			start=start+posIndex+1
// 		}
// 	}

// 	return maxLen
// }

func mySqrt(x int) int {
    left:=float64(0)
    right:=float64(x)        

    for {
		mid:=float64((left+right)/2)
        res:=mid*mid
        if(res==float64(x)) {
            return int(mid)
        } else if(res>float64(x)) {
            right=mid
        } else {
            left=mid
        }
        
        if(int(left)==int(right)) {
            return int(left)
		}
		
    }
    
    return int(left)
}