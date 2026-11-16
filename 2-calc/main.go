package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	for {

		var valueOfUser string
		arrNumbers := []float64{}
		result := 0.0

		typeOperationUser, err := getTypeOperation()
		if err != nil {
			fmt.Println(err)
			continue
		}
		valueOfUser, err = getValueOfUser()
		if err != nil {
			fmt.Println(err)
			continue
		}

		arrNumbers, err = convertStrToSlice(valueOfUser)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch typeOperationUser {
		case "AVG":
			lenArr := float64(len(arrNumbers))
			for _, value := range arrNumbers {
				result += value
			}
			result = result / lenArr
			fmt.Printf("%s = %.2f\n", typeOperationUser, result)
			return
		case "SUM":
			for _, value := range arrNumbers {
				result += value
			}
			fmt.Printf("%s = %.2f\n", typeOperationUser, result)
			return
		case "MED":
			result = medianFloat(arrNumbers)
			fmt.Printf("%s = %.2f\n", typeOperationUser, result)
			return
		default:
			fmt.Println("Something went wrong")
			return
		}
	}
}

func convertStrToSlice(values string) ([]float64, error) {
	line := strings.Split(values, ",")
	nums := make([]float64, 0, len(line))

	for _, value := range line {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, errors.New("Not all items in the row are numbers")
		}
		nums = append(nums, num)
	}
	return nums, nil

}

func getValueOfUser() (string, error) {
	var values string
	fmt.Println("Enter the numbers separated by commas without space")
	n, err := fmt.Scan(&values)
	if err != nil {
		var dump string
		fmt.Scanln(&dump)
		return values, errors.New("invalid input")
	}
	if n != 1 {
		return values, errors.New("empty input")
	}
	return values, nil

}

func getTypeOperation() (string, error) {
	var typeOperaion string
	fmt.Println("Enter type operation (AVG, SUM, MED)")
	_, err := fmt.Scan(&typeOperaion)
	if err != nil {
		var dump string
		fmt.Scanln(&dump)
		return "", errors.New("invalid input")
	}
	typeOperaionToUpper := strings.ToUpper(typeOperaion)

	if typeOperaionToUpper == "AVG" || typeOperaionToUpper == "SUM" || typeOperaionToUpper == "MED" {
		return typeOperaionToUpper, nil
	}
	return typeOperaionToUpper, errors.New("Enter type AVG or SUM or MED ")

}

func medianFloat(nums []float64) float64 {

	sorted := slices.Clone(nums)
	slices.Sort(sorted)

	n := len(sorted)
	mid := n / 2
	if n%2 == 1 {
		return float64(sorted[mid])
	}
	return float64(sorted[mid-1]+sorted[mid]) / 2.0
}
