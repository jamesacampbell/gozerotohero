package main

import (
	"fmt"
)

func main() {
	people := [3]string{"helen", "ark", "dick"}
	friends := [2]string{"ark", "dick"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction.")
	i := 0
loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}
	language := "golang"
	switch language {
	case "Python":
		fmt.Println("You are learning Python")
	}

}
