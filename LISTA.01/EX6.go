package main

import "fmt"

var qnt_temp int
var temp, temp_Celsius float64
var Temp_Fharenheit []float64
var Temp_Celsius []float64

func main() {
	fmt.Print("Quantas temperaturas vão ser convertidas? ")
	fmt.Scan(&qnt_temp)

	for i := 0 ; i < qnt_temp ; i++ {
		fmt.Println("Escreva a temperatura",i + 1,"em fharenheit: ")
		fmt.Scan(&temp)
		Temp_Fharenheit = append(Temp_Fharenheit, temp)
		converterTemp()
		Temp_Celsius = append(Temp_Celsius, temp_Celsius)
	}

	for i := 0 ; i < qnt_temp; i++ {
		fmt.Printf("%.2f FHARENHEIT EQUIVALE A %.2f CELSIUS\n", Temp_Fharenheit[i], Temp_Celsius[i])
	}
}
func converterTemp() {
	temp_Celsius = 5*(temp - 32) / 9
}
