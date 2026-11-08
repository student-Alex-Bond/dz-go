package main

import "fmt"

func main() {
	
	const USDtoEUR float64 = 0.8909
	const USDtoRUB float64 = 82.14
	const EURtoRUB float64 = USDtoRUB / USDtoEUR 
	fmt.Printf("EUR to RUB = %.2f rub\n" , EURtoRUB)
}