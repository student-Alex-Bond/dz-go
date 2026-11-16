package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	for {

		valueOfUser := bufio.NewScanner(os.Stdin)
		arrNumbers := []float64{}
		result := 0.0

		typeOperationUser, err := getTypeOperation(valueOfUser)
		if err != nil {
			fmt.Println(err)
			continue
		}
		lineOfNumbers, err := getValueOfUser(valueOfUser)
		if err != nil {
			fmt.Println(err)
			continue
		}

		arrNumbers, err = convertStrToSlice(lineOfNumbers)
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
	if len(nums) == 0 {
		return nil, errors.New("Empty line of numbers")
	}
	return nums, nil

}

func getValueOfUser(scan *bufio.Scanner) (string, error) {
	fmt.Println("Enter the numbers separated by commas without space")
	if !scan.Scan() {
		return "", errors.New("input break")
	}
	return scan.Text(), nil

}

func getTypeOperation(scan *bufio.Scanner) (string, error) {
	fmt.Println("Enter type operation (AVG, SUM, MED)")
	if !scan.Scan() {
		return "", errors.New("ввод прерван")
	}
	typeOperaionToUpper := strings.ToUpper(scan.Text())

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
