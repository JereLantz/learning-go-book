package main

import "fmt"

// Solutions https://github.com/learning-go-book-2e/ch08

/* 1.
Write a generic function that doubles the value of any integer or float that’s passed in to it.
Define any needed generic interfaces.
*/
type Number interface{
	~int | ~int8 | ~int16 | ~int32 | ~int64 | 
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

func doubler[N Number] (n N) any{
	return n * 2
}

/* 2.
Define a generic interface called Printable that matches a type that implements
fmt.Stringer and has an underlying type of int or float64.
Define types that meet this interface.
Write a function that takes in a Printable and prints its value to the screen
using fmt.Println.
*/
type Printable interface{
	fmt.Stringer
	~int | ~float64
}

func (i CustomInt) String() string {
	return fmt.Sprintf("{%d}", i)
}

type CustomInt int

func customPrint[t Printable](message t){
	fmt.Println(message)
}

/* 3.
Write a generic singly linked list data type.
Each element can hold a comparable value and has a pointer to the next element in the list.
The methods to implement are as follows:

// adds a new element to the end of the linked list
Add(T)
// adds an element at the specified position in the linked list
Insert(T, int)
// returns the position of the supplied value, -1 if it's not present
Index (T) int
*/
type node[T comparable] struct {
	val T
	next *node[T]
}

type list[T comparable] struct {
	head *node[T]
	tail *node[T] 
}

func (l *list[T]) Add(t T){
	// oma yritys. Toimii mutta siihen oli parempi ratkaisu
	if l.head == nil {
		l.head = &node[T]{val:t}
		l.tail = l.head
	}else {
		current := l.head
		for current.next != nil {
			current = current.next
		}

		current.next = &node[T]{val:t}
		l.tail = current.next
	}
	/* Parempi ratkaisu olisi ollut:

		n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
	*/
}

func (l *list[T]) Index(t T) int{
	current := l.head
	i := 0
	for ; current.val != t ; i++{
		if current.next == nil {
			return -1
		}
		current = current.next
	}
	return i

	/* Kirjan kirjoittajan ratkaisu
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
	*/
}

func (l *list[T]) Insert(t T, index int){
	n := &node[T]{val:t,}

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	if index == 0 {
		oldHead := l.head
		l.head = n
		n.next = oldHead
		return
	}

	befInsertNode := l.head

	for range index-1{
		if befInsertNode.next == nil {
			return
		}
		befInsertNode = befInsertNode.next
	}

	if befInsertNode.next == nil {
		l.Add(t)
		return
	}

	befInsertNode.next = &node[T]{val:t, next: befInsertNode.next}
}

func main(){
	num := 3.3
	doubledNum := doubler(num)

	fmt.Println("num:" , num, "double:", doubledNum)

	var x CustomInt = 69

	customPrint(x)

	fmt.Println("========Ex 3========")
	l := &list[int]{}
	l.Add(69)
	l.Add(50)
	l.Add(30)


	node := l.head
	for node.next != nil {
		fmt.Println(node.val)
		node = node.next
	}
	fmt.Println(node.val)


	fmt.Println("Tail:", l.tail.val)
	fmt.Println("Head:", l.head.val)

	find := 69
	fmt.Println("Index of:", find, "=", l.Index(find))


	l.Insert(1, 1)
	l.Insert(3, 0)

	node = l.head
	for node.next != nil {
		fmt.Println(node.val)
		node = node.next
	}
	fmt.Println(node.val)


	fmt.Println("Tail:", l.tail.val)
	fmt.Println("Head:", l.head.val)
}
