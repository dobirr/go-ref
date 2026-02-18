package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Geben Sie die Länge des Quaders ein:")
	fmt.Scan(&a)

	fmt.Println("Geben Sie die Breite des Quaders ein:")
	fmt.Scan(&b)

	fmt.Println("Geben Sie die Tiefe des Quaders ein:")
	fmt.Scan(&c)

	fmt.Printf("Ein %v x %v x %v Quader hat die geometrischen Größen:\n", a, b, c)
	fmt.Printf("Volumen: %v\n", a*b*c)
	fmt.Printf("Kantensumme: %v\n", a*4+b*4+c*4)
	fmt.Printf("Oberfläche: %v\n", 2*(a*b+a*c+b*c))
	fmt.Printf("Umkugelradius: %.2f\n", math.Sqrt(a*a+b*b+c*c)/2)
	fmt.Printf("Raumdiagonale: %.2f\n", math.Sqrt(a*a+b*b+c*c))

	fmt.Println("")
}
