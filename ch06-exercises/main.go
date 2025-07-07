package main

import "fmt"

/* 1.
Create a struct named Person with three fields:
FirstName and LastName of type string and Age of type int.
Write a function called MakePerson that takes in firstName, lastName, and age and returns a Person.
Write a second function MakePersonPointer that takes in firstName,
lastName, and age and returns a *Person.
Call both from main.
Compile your program with go build -gcflags="-m".
This both compiles your code and prints out which values escape to the heap.
Are you surprised about what escapes?
*/

type Person struct{
	FirstName string
	LastName string
	Age int
}

func MakePerson(fName, lName string, age int) Person{
	return Person{
		FirstName: fName,
		LastName: lName,
		Age: age,
	}
}

func MakePersonPointer(fName, lName string, age int) *Person{
	return &Person{
		FirstName: fName,
		LastName: lName,
		Age: age,
	}
}

func ex1(){
	p1 := MakePerson("Etu", "Suku", 69)
	p2 := MakePersonPointer("Jorma", "Jormala", 69)

	fmt.Println(p1)
	fmt.Println(p2)
}

/* 2.
Write two functions.
The UpdateSlice function takes in a []string and a string.
It sets the last position in the passed-in slice to the passed-in string.
At the end of UpdateSlice, print the slice after making the change.

The GrowSlice function also takes in a []string and a string.
It appends the string onto the slice.
At the end of GrowSlice, print the slice after making the change.

Call these functions from main.
Print out the slice before each function is called and after each function is called.
Do you understand why some changes are visible in main and why some changes are not?
*/
func GrowSlice(slice []string, str string){
	slice = append(slice, str)
	fmt.Println(slice)
}

func UpdateSlice(slice []string, str string){
	slice[len(slice)-1] = str
	fmt.Println(slice)
}

func ex2(){
	slice := []string{"a", "b", "c", "d"}

	fmt.Println("main", slice)
	UpdateSlice(slice, "e")

	fmt.Println("main", slice)
	GrowSlice(slice, "g")

	fmt.Println("main", slice)
}

/* 3.
Write a program that builds a []Person with 10,000,000 entries
(they could all be the same names and ages).
See how long it takes to run.
Change the value of GOGC and see how that affects the time it takes for the program to complete.
Set the environment variable GODEBUG=gctrace=1 to see when garbage collections happen
and see how changing GOGC changes the number of garbage collections.
What happens if you create the slice with a capacity of 10,000,000?
*/
func main(){
	ex1()
	ex2()
}
