package main

import (
	"fmt"
	"math/rand"
)

/*
3.
Start a new program.
In main, declare an int variable called total.
Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
The body of the for loop should be as follows:

	total := total + i
	fmt.Println(total)

After the for loop, print out the value of total.
What is printed out? What is the likely bug in this code?
*/
func ex3() {
	total := 0

	for i := range 10 {
		total = total + i
	}
	fmt.Println(total)
}

/*
2.
Loop over the slice you created in exercise 1.
For each value in the slice, apply the following rules:

	a.If the value is divisible by 2, print “Two!”
	b.If the value is divisible by 3, print “Three!”
	c.IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
	d.Otherwise, print “Never mind”.
*/
func ex2(numbers []int) {
	for _, v := range numbers{
		switch {
		case v % 2 == 0 && v % 3 == 0:
			fmt.Println(v, "Six!")

		case v % 3 == 0:
			fmt.Println(v, "Three!")

		case v % 2 == 0:
			fmt.Println(v, "Two!")

		default:
			fmt.Println(v, "Never mind")
		}
	}
}

func main() {
	/*
	1.
	Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
	*/
	numbers := make([]int, 0, 100)

	for range 100 {
		numbers = append(numbers, rand.Intn(101))
	}
	fmt.Println(numbers)

	ex2(numbers)
	ex3()
}
