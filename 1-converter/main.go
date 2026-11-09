package main

import (
	"fmt"
)

func main() {
	value := getUserInput()
	fmt.Println(value)
	const USDtoEUR float64 = 0.8909
	const USDtoRUB float64 = 82.14
	const EURtoRUB float64 = USDtoRUB / USDtoEUR
	fmt.Printf("EUR to RUB = %.2f rub\n", EURtoRUB)
}

func calculationCurrency(sumOfMoney float64, sourceCurrency, targetCurrency string) float64 {
	return 0
}

func getUserInput() string {
	var value string
	fmt.Print("Введите сумму: ")
	fmt.Scan(&value)
	return value
}
