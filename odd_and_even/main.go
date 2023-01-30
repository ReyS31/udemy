package main

import "fmt"

func main() {
	var list []int

	for i := 0; i <= 10; i++ {
		list = append(list, i)
	}

	for i := range list {
		if i%2 == 0 {
			fmt.Println(i, "is", "even")
		} else {
			fmt.Println(i, "is", "odd")
		}
	}
}
