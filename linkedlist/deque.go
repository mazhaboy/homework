package main

import (
	"fmt"
)
//implement DEQUEUE
type Deque struct {
	front *Node
	back *Node
}
type Node struct {
	value interface{}
	next *Node
}


func main() {
	dequeue:=NewDeque()
	dequeue.PushFront(3)
	dequeue.PushFront(2)
	dequeue.PushFront(1)
	dequeue.PushBack(4)
	dequeue.PushBack(5)
	dequeue.PushBack(6)
	dequeue.PrintDeque()
	fmt.Println(dequeue.PopFront())
	fmt.Println(dequeue.PopBack())
	dequeue.PrintDeque()
	dequeue.PushFront("start")
	dequeue.PushBack("finish")
	dequeue.PrintDeque()
	dequeue.Peek(true)
	dequeue.Peek(false)


}
func(deque *Deque)PrintDeque(){
	l:=deque.front
	for l!=nil{
		fmt.Println("value: ",l.value)
		l=l.next
	}
}
func NewDeque()*Deque{
	return new(Deque)
}
func (deque *Deque)PushFront(value interface{}){
	new_node:=&Node{
		value: value,
		next:nil,
	}
	if deque.front==nil && deque.back==nil{
		deque.front=new_node
		deque.back=new_node
	} else {
		x:=deque.front
		deque.front=new_node
		deque.front.next=x
	}
}

func (deque *Deque)PushBack(value interface{}){
	new_node:=&Node{
		value: value,
		next:nil,
	}
	if deque.front==nil && deque.back==nil{
		deque.front=new_node
		deque.back=new_node
	} else {
		deque.back.next=new_node
		deque.back=new_node
	}
}

func (deque *Deque)PopFront()interface{}{

	if deque.front==nil && deque.back==nil{
		return nil
	} else {
		x:=deque.front.next
		deleted:=deque.front.value
		deque.front=nil
		deque.front=x
		return deleted
	}
}

func (deque *Deque)PopBack()interface{}{
	if deque.front==nil && deque.back==nil{
		return nil
	} else {
		y:=deque.front
		x:=deque.back
		for y.next!=x{
			y=y.next
		}
		deque.back=nil
		deque.back=y
		deque.back.next=nil
		return x.value
	}
}

func (deque *Deque)Peek(IsFront bool){
	if deque.front!=nil && IsFront {
		fmt.Println("value: ", deque.front.value)
	} else if deque.back!=nil && !IsFront {
		fmt.Println("value: ", deque.back.value)
	} else {
		fmt.Println("value: ", "null")
	}
}

//value:  1
//value:  2
//value:  3
//value:  4
//value:  5
//value:  6
//1
//6
//value:  2
//value:  3
//value:  4
//value:  5
//value:  start
//value:  2
//value:  3
//value:  4
//value:  5
//value:  finish
//value:  start
//value:  finish
