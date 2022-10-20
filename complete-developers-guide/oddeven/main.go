package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range numbers {
		t := "odd"
		if n%2 == 0 {
			t = "even"
		}
		fmt.Printf("%v is %v\n", n, t)
	}
}
