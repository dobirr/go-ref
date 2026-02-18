package main

import (
	"fmt"
)

func main() {

	var sizeM float64

	fmt.Println("Geben Sie eine Länge (in Metern) ein:")
	fmt.Scan(&sizeM)

	fmt.Println("Geben Sie eine Länge (in Metern) ein:")

	fmt.Printf("Meter entsprechen: %v m. \n", sizeM)
	fmt.Printf("Millimeter entsprechen: %e mm. \n", sizeM*1000)
	fmt.Printf("Kilometer entsprechen: %v km. \n", sizeM/1000)
	fmt.Printf("Zoll entsprechen: %6.3f zoll. \n", sizeM*100/2.54)
	fmt.Printf("Seemeilen entsprechen: %4.2f sm. \n", sizeM/1852)
	fmt.Printf("Lichtjahre entsprechen: %e Lj. \n", sizeM/9_460_730_472_580_800)

}
