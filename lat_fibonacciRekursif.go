package main

import (
	"fmt"
)

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	var bil int = 10
	//var arrBil [...]int
	fmt.Printf("Fibonacci suku ke-%d : %d\n", bil, fibo(bil))
	fmt.Printf("Deret fibonacci sampai suku ke-%d\n", bil)
	for i := 0; i <= bil; i++ {
		fmt.Printf("%d ", fibo(i))
	}
}
