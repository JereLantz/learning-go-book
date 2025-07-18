package main

import (
	"fmt"
	"sync"
)

/*
Create a function that launches three goroutines that communicate using a channel.
The first two goroutines each write 10 numbers to the channel.
The third goroutine reads all the numbers from the channel and prints them out.
The function should exit when all values have been printed out.
Make sure that none of the goroutines leak.
You can create additional goroutines if needed.
*/
func process(){
	var wg sync.WaitGroup
	routines := 2

	wg.Add(routines)

	ch := make(chan int)

	for i := range routines {
		go func(i int){
			defer wg.Done()
			for v := range 10 {
				ch <- i*10+v
			}
		}(i)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)

	go func(){
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg2.Wait()
}

func ProcessSolution() {
	// copied from the solution
	ch := make(chan int)
	// use 2 waitgroups!
	// the 1st waitgroup controls when to close the channel
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i*100 + 1
		}
	}()
	// launch this helper goroutine to close the channel when the two writing goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()
	// the second waitgroup signals when the reading goroutine is done
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg2.Wait()
}

func main(){
	process()
	//ProcessSolution()
}
