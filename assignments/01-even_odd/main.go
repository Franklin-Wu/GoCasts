package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range slice {
		if v % 2 == 0 {
			fmt.Printf("%v is even\n", v)
		} else {
			fmt.Printf("%v is odd\n", v)
		}
	}
}
