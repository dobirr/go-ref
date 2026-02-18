package main

import (
	"fmt"
)

func main() {
	var hours, minutes int
	var hoursAngle, minutesAngel float64

	fmt.Println("Geben Sie bitte eine Uhrzeit an, indem Sie zunächst\ndie Stunde (von 0 bis 23) und dann die Minute (von 0 bis 59) eingeben:")
	fmt.Scan(&hours, &minutes)

	if hours > 11 {
		hours = hours - 12
	}

	hoursAngle = 360.0/12.0*float64(hours) + 360.0/(12*60.0)*float64(minutes)
	minutesAngel = 360.0 / 60.0 * float64(minutes)

	fmt.Printf("Zeigerstellung um %02d:%02d Uhr:\n", hours, minutes)
	fmt.Printf("Winkel des Stundenzeigers: %.1f°\n", hoursAngle)
	fmt.Printf("Winkel des Minutenzeigers: %.1f°\n", minutesAngel)
}
