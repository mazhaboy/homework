package main

import (
	"fmt"
)
func main(){
 fmt.Println(GetMinAndMax(0,7,99,2))
}
func GetMinAndMax(a,b,c,d int) (int,int) {
	arr:= []int{}
	arr=append(arr,a,b,c,d)
	BubbleSort(&arr)
	max:=len(arr)-1
	return arr[0],arr[max]
}
func BubbleSort(arr *[]int){
	for i:=0; i<len(*arr)-1; {
		first:=(*arr)[i]
		second:=(*arr)[i+1]
		if (*arr)[i]>(*arr)[i+1]{
			(*arr)[i]=second
			(*arr)[i+1]=first
			i=0
		} else {
			i++
		}
	}
}