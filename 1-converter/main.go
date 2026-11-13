package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	currentCurrency := getUserInputCurrentCurrency()
	targetCurrency := getUserInputTargetCurrency(currentCurrency)
	currencyAmount := getUserInputCurrencyAmount()
	rateUSDtoEUR := 0.92
	rateUSDtoRUB := 92.5
	rateEURtoRUB := 100.5
	result := 0.0
	switch {
	case currentCurrency == "USD" && targetCurrency == "EUR":
		result = currencyAmount * rateUSDtoEUR
	case currentCurrency == "USD" && targetCurrency == "RUB":
		result = currencyAmount * rateUSDtoRUB
	case currentCurrency == "EUR" && targetCurrency == "USD":
		result = currencyAmount / rateUSDtoEUR
	case currentCurrency == "EUR" && targetCurrency == "RUB":
		result = currencyAmount * rateEURtoRUB
	case currentCurrency == "RUB" && targetCurrency == "USD":
		result = currencyAmount / rateUSDtoRUB
	case currentCurrency == "RUB" && targetCurrency == "EUR":
		result = currencyAmount / rateEURtoRUB
	default:
		fmt.Println("Неподдерживаемая пара валют")
		return
	}

	fmt.Println("Исходная валюта", currentCurrency)
	fmt.Println("Сумма валюты", currencyAmount)
	fmt.Println("Целевая валюта", targetCurrency)
	fmt.Printf("%s => %s = %.2f %s\n", currentCurrency, targetCurrency, result, targetCurrency)
}

func getUserInputCurrentCurrency() string {
	var currentCurrency string
	for {
		fmt.Println("Введите исходнную валюту (USD, EUR, RUB)")
		fmt.Scan(&currentCurrency)
		currentCurrencyToUpper := strings.ToUpper(currentCurrency)
		if currentCurrencyToUpper == "USD" || currentCurrencyToUpper == "EUR" || currentCurrencyToUpper == "RUB" {
			return currentCurrencyToUpper
		}
		fmt.Println("Выберите из имеющихся USD, EUR, RUB")
	}
}
func getUserInputTargetCurrency(currentCurrency string) string {
	var targetCurrency string

	for {
		switch currentCurrency {
		case "USD":
			fmt.Println("Введите целевую валюту ( EUR, RUB)")
			fmt.Scan(&targetCurrency)
			targetCurrencyToUpper := strings.ToUpper(targetCurrency)
			fmt.Println(targetCurrencyToUpper)
			if targetCurrencyToUpper == "EUR" || targetCurrencyToUpper == "RUB" {
				return targetCurrencyToUpper
			}
			fmt.Println("Выберите из имеющихся  EUR, RUB")
		case "EUR":
			fmt.Println("Введите целевую валюту ( USD, RUB)")
			fmt.Scan(&targetCurrency)
			targetCurrencyToUpper := strings.ToUpper(targetCurrency)
			if targetCurrencyToUpper == "USD" || targetCurrencyToUpper == "RUB" {
				return targetCurrencyToUpper
			}
			fmt.Println("Выберите из имеющихся  USD, RUB")
		case "RUB":
			fmt.Println("Введите целевую валюту ( EUR, USD)")
			fmt.Scan(&targetCurrency)
			targetCurrencyToUpper := strings.ToUpper(targetCurrency)
			if targetCurrencyToUpper == "EUR" || targetCurrencyToUpper == "USD" {
				return targetCurrencyToUpper
			}
			fmt.Println("Выберите из имеющихся  EUR, USD")
		}

	}
}
func getUserInputCurrencyAmount() float64 {
	var currencyAmount string
	for {
		fmt.Println("Введите сумму валюты")
		fmt.Scan(&currencyAmount)
		countCurrency, err := strconv.ParseFloat(currencyAmount, 64)
		if err != nil && countCurrency <= 0 {
			fmt.Println("Не число или число меньше нуля, попробуйте ещё раз")
			continue
		}
		return countCurrency
	}

}
