package main

import (
	"fmt"
	"strconv"
	"strings"
)
var rates = map[string]map[string]float64{
	"USD": {"EUR": 0.92, "RUB": 92.5},
	"EUR": {"USD": 1 / 0.92, "RUB": 100.5},
	"RUB": {"USD": 1 / 92.5, "EUR": 1 / 100.5},
}

func main() {
	currentCurrency := getUserInputCurrentCurrency()
	currencyAmount := getUserInputCurrencyAmount()
	targetCurrency := getUserInputTargetCurrency(currentCurrency)
	result, err := convert(currencyAmount, currentCurrency, targetCurrency)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Исходная валюта", currentCurrency)
	fmt.Println("Сумма валюты", currencyAmount)
	fmt.Println("Целевая валюта", targetCurrency)
	fmt.Printf("%s => %s = %.2f %s\n", currentCurrency, targetCurrency, result, targetCurrency)
}

func convert(amount float64, from, to string) (float64, error) {
	rate, ok := rates[from][to] // второй аргумент: есть ли пара
	if !ok {
		return 0, fmt.Errorf("неподдерживаемая пара валют %s -> %s", from, to)
	}
	return amount * rate, nil
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
