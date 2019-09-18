package main

import "fmt"

func main() {

	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for oddOrEven := range intSlice {
		if oddOrEven%2 != 0 {
			fmt.Printf("%v is odd\n", oddOrEven)
		} else {
			fmt.Printf("%v is even\n", oddOrEven)
		}
	}

}
