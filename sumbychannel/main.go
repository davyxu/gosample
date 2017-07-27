package main

import (
	"fmt"
)

func sum(n int, c chan int) {

	var result int
	for i := 0; i <= n; i++ {
		result += i
	}
	c <- result
}

func main() {

	c := make(chan int)

	go sum(100, c)

	fmt.Println(<-c)
}
