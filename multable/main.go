package main

import "fmt"

func main() {

	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d*%d=%d ", y, x, y*x)
		}

		fmt.Println()
	}
}
