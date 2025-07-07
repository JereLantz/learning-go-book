package main

import (
	"time"
	"fmt"
)


type Counter struct {
    total       int
    lastUpdated time.Time
}

func (c *Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}

func (c Counter) String() string {
    return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
    c.Increment()
    fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
    c.Increment()
    fmt.Println("in doUpdateRight:", c.String())
}

type joku interface{
	yksi()
	kaksi()
}

func (jokuStruct)yksi(){
}
func (jokuStruct)kaksi(){
}

type jokuStruct struct{
	jee string
	jeejee int
}

func jokuFn(j joku){
}

func main() {
	var j jokuStruct
	jokuFn(j)

    var c Counter
    doUpdateWrong(c)
    fmt.Println("in main:", c.String())
    doUpdateRight(&c)
    fmt.Println("in main:", c.String())
}
