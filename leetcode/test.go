package main

import (
	"fmt"
	"time"
)

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
  )

  const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)
  
  var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
  )
  
  func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
	  return fmt.Sprintf(colorString,
		fmt.Sprint(args...))
	}
	return sprint
  }

func main() {
	//var arr=[]int{2,3,6,4,7}
	//var target=9

	start_time:=time.Now()

	//color.Red("*")
	//fmt.Println(twoSum(arr,target))
	DrawHeart()

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

func DrawHeart() {
	for y := 1.4; y > -1.4; y -= 0.1 {
        for x := -1.4; x < 1.4; x += 0.04 {
			a := x * x + y * y - 1;
			if a * a * a - x * x * y * y * y <= 0.0 {
				fmt.Printf("☆")
			} else {
			fmt.Printf(" ")
			}
        }
        fmt.Println("")
    }
}