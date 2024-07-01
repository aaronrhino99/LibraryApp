package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
  prev *Node[T]
  value T
}

type LinkedList[T any] struct {
  head   *Node[T]
  length uint
}

func (ll *LinkedList[T]) index(idx uint) ( T ,bool) {
  
  var value T 
  if idx >=ll.length {
    return value , false
  }

  currentIndex := 0
  currentValue := ll.head
  for currentValue != ll.head || currentIndex == 0 {
    if uint(currentIndex) == idx {
      value = currentValue.value
      break
    }
    currentValue = currentValue.next
    currentIndex += 1
  }

  return value, true
 
}

func (ll *LinkedList[T]) append(value T) {
 ll.length++
 if ll.length == 1 {
  ll.head = &Node[T]{}
  ll.head.prev = ll.head
  ll.head.next = ll.head
  ll.head.value = value
 }

 node := &Node[T]{}
 node.value = value

 tail := ll.head.prev
 tail.next = node
 node.prev = tail
 node.next = ll.head
 ll.head.prev = node
}

//func (ll *LinkedList[T]) pop() T {
  
//}

func (ll *LinkedList[T]) printList() {
  values := [T]{}

  currentValue = ll.head// it dereferences by its self remember that 
  for ll.length != 0{
    values = append(values, currentValue.value)
    currentValue = currentValue.next
    if currentValue == ll.head {
      break
    }
  }
 fmt.Println(values)
}

func main() {
  ll := LinkedList[int]{}
  ll.append(123)
  ll.append(1)
  ll.append(23)
  ll.append(3)
  ll.printList()
}