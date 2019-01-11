package main

import "fmt"

func main() {

	x := 0

	for {
		if x <= 10000 {
			fmt.Println(x)
			x++
		}
	}
}
