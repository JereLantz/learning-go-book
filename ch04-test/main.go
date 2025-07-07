package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// go:ssa pystytään luomaan muuttujia joiden scope on if:n sisällä, luomalla
	// se ennen vertausta.

	if n := rand.IntN(10); n == 0 {
		fmt.Println("Too low")
	} else if n > 5 {
		fmt.Println("that is too big:", n)
	} else {
		fmt.Println("Tha's a good number:", n)
	}

	// Tämä aiheuttaa compile error
	//fmt.Println(n)

	//Go myös shadowaa muttujia
	x := 5
	if x > 0 {
		fmt.Println(x)
		x := 10
		fmt.Println(x)
	}
	fmt.Println(x)

	// shadowaamisen kanssa pitää olla varovainen välillä, ettei vahingossa
	// muokkaa/ole muokkaamatta jotain mitä olisi pitänyt
	y := 10
	if y > 0 {
		fmt.Println(y)
		// tämä ei muokkaa y:tä, vaan shadowaa sen tämän if block:in ajaksi
		y, z := 20,30
		fmt.Println(y, z)
	}
	fmt.Println(y)

	fmt.Println("For range")
	for i := range 10 {
		fmt.Println(i)
	}
}
