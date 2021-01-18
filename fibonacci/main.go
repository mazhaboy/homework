package main

import "fmt"

func main(){
	fmt.Println(Fibonacci(6))
}
func Fibonacci(index int) int {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else if index > 1 {
		return Fibonacci(index-1) + Fibonacci(index-2)
	} else {
		return -1
	}
}