package main

import "fmt"

func power(a, b int) int {
	for i := 1; i <=b; i++ {
		a *= b
	}
	return a
}

func main()  {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	fmt.Print(power(a, b))
}
