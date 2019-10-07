package main

import "fmt"
import "math"
import "time"

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("It is the Weekend")
	default :
		fmt.Println("It is the Weekday")
	fmt.Println(time.Now().Weekday())	
	fmt.Println(time.Now())
	fmt.Println(time.Now().Hour())
	}
}