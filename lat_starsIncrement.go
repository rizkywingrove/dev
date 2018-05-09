package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 6; i++ {
		for j := 0; j < i; j++ {
			if i == 4 {
				break
			}

			//ini komen
			fmt.Print("*", " ")
		}

		if i != 4 {
			fmt.Println()
		}

	}
}
