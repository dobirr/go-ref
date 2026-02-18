package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Println("Geben Sie die drei Längen des Dreiecks ein")
	fmt.Scan(&a, &b, &c)

	istRechtwinklig(a, b, c)

	fmt.Println("ist rechtwinkelig:", istRechtwinklig(a, b, c))
}
