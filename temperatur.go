package main

import "fmt"

func convertToReamur(celciusFormat float32) float32 {
	var reamurValue float32 = celciusFormat + 273.0
	return reamurValue

}

func main() {
	var temperaturCelcius float32

	fmt.Scanln(temperaturCelcius)
	var hasil float32 = convertToReamur(temperaturCelcius)
	fmt.Print(hasil)

}
