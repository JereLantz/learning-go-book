package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/* 3.
Write a function called prefixer that has an input parameter of type string and
returns a function that has an input parameter of type string and returns a string.
The returned function should prefix its input with the string passed into prefixer.
Use the following main function to test prefixer:

func main() {
    helloPrefix := prefixer("Hello")
    fmt.Println(helloPrefix("Bob")) // should print Hello Bob
    fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
*/
func ex3(){
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob")) // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

func prefixer(s string) (func(s1 string) string){
	return func(s1 string) string {
		return s + " " + s1
	}
}

/* 2.
Write a function called fileLen that has an input parameter of type string and
returns an int and an error.
The function takes in a filename and returns the number of bytes in the file.
If there is an error reading the file, return the error.
Use defer to make sure the file is closed properly.
*/
func ex2(){
	if len(os.Args) < 2{
		log.Println("no filename specified")
		return
	}
	fl, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File length is", fl)
}

func fileLen(filename string) (int, error){
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data := make([]byte,2048)
	total := 0
	for {
		count, err := file.Read(data)
		total += count
		if err != nil {
			if err != io.EOF{
				return 0, err
			}
			break
		}
	}

	return total, nil
}

/* 1.
The simple calculator program doesn’t handle one error case: division by zero.
Change the function signature for the math operations to return both an int and an error.
In the div function, if the divisor is 0, return errors.New("division by zero") for the error.
In all other cases, return nil.
Adjust the main function to check for this error.
*/
func ex1(){
}

func main(){
	ex1()
	ex2()
	ex3()
}
