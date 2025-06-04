package main

import "fmt"

func main() {
	sum := 0
	maxLoop := 10

	for i := maxLoop; i > 0; i-- {
		fmt.Println(i)
	}

	for range maxLoop {
		sum += 1
		fmt.Println(sum)
	}

}