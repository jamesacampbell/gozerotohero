package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}
}
