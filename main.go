package main

import "fmt"

func main() {
	isEven(2)
}
func isEven(n int) {
	var result string = "is not even!"
	if n%2 == 0 {
		result = "is even!"
	}
	fmt.Println(result)
}
