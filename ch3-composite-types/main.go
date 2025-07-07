package main

import "fmt"

/*
1.
Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
Create a subslice containing the first two values; a second subslice with the second, third, and fourth values; and a third subslice with the fourth and fifth values.
Print out all four slices.
*/
func firstEx(){
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт",
	}

	first2Greet := greetings[:2]
	fmt.Println(first2Greet)
	secondThirdFourthGreet := greetings[1:4]
	fmt.Println(secondThirdFourthGreet)
	fourthFifthGreet := greetings[3:]
	fmt.Println(fourthFifthGreet)
}
/*
2.
Write a program that defines a string variable called message with the value "Hi and " and prints the fourth rune in it as a character, not a number.
*/
func secondEx(){
	message := "Hi 🐒 and 🧑🏼‍🦯"
	fmt.Println(message)

	msgRune := []rune(message)

	fmt.Println(string(msgRune[3]))
}

/*
3.
Write a program that defines a struct called Employee with three fields:
firstName, lastName, and id.
The first two fields are of type string, and the last field (id) is of type int.
Create three instances of this struct using whatever values you’d like.
Initialize the first one using the struct literal style without names,
the second using the struct literal style with names, and the third with a var declaration.
Use dot notation to populate the fields in the third struct.
Print out all three structs.
*/
func thirdEx(){
	type employee struct{
		firstName string
		lastName string
		id int
	}

	employee1 := employee{
		"bob",
		"theBuilder",
		69,
	}
	employee2 := employee{
		firstName: "John",
		lastName: "Go",
		id: 3,
	}
	var employee3 = employee{}
	employee3.id = 4
	employee3.firstName = "John"
	employee3.lastName = "Rust"

	fmt.Println("Employee 1: ", employee1)
	fmt.Println("Employee 2: ", employee2)
	fmt.Println("Employee 3: ", employee3)
}

func main(){
	firstEx()
	secondEx()
	thirdEx()
}
