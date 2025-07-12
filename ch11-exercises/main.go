package main

import (
	"embed"
	"fmt"
	"os"
)

/*
Go to the UN’s Universal Declaration of Human Rights (UDHR) page and copy the
text of the UDHR into a text file called english_rights.txt.
Click the Other Languages link and copy the document text in a few additional languages into
files named LANGUAGE_rights.txt.
Create a program that embeds these files into package-level variables.
Your program should take in one command-line parameter, the name of a language.
It should then print out the UDHR in that language.
*/

//go:embed rights/*
var Rights embed.FS

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Please insert language")
		return
	}
	fmt.Println(os.Args[1])
	data, err := Rights.ReadFile("rights/" + os.Args[1] + "_rights.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
