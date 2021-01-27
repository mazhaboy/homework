package main

import (
	"fmt"
)
func main(){
	fmt.Println(Add(2,3))
	fmt.Println(Sum(3))
	fmt.Println(Lifehack(0))
	fmt.Println(Lifehack(5))
	fmt.Println(Lifehack(75))
	fmt.Println(Lifehack(4255))
}
func Add(a, b int) int {
	return a+b
}
func Sum(a int) int {
	x:=0
	for i:=0; i<=a;i++{
		x=x+i
	}
	return x
}
func Lifehack(a int) int {
	if a==0 {
		return 0
	}
	x:=a/10
	result:=x*(x+1)
	return(result*100)+25
}
