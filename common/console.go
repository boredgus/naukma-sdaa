package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScanNumber(reader *bufio.Reader, msg string, minimal *int) int {
	for {
		fmt.Print(msg)
		amountStr := readInput(reader)
		value, err := strconv.Atoi(amountStr)
		if err != nil {
			fmt.Println("Invalid data provided. Enter a number.", err)
			continue
		}

		if minimal != nil && value < *minimal {
			fmt.Printf("Invalid data provided. Value should be %v or larger.\n", minimal)
			continue
		}

		return value
	}
}

func ScanUniqueNumbers(reader *bufio.Reader, amount int) []int {
	for {
		fmt.Printf("Enter %v numbers separated with 1 whitespace: ", amount)
		inputStr := readInput(reader)

		stringifiedNums := strings.Split(inputStr, " ")
		if len(stringifiedNums) != amount {
			fmt.Printf("Invalid amount of numbers privided. %v required.\n", amount)
			continue
		}

		nums, err := parseNumbers(stringifiedNums)
		if err != nil {
			fmt.Println("Invalid data provided: some of elements is not a number")
			continue
		}

		if hasDuplicates(nums) {
			fmt.Println("Invalid data provided: no duplicates allowed")
			continue
		}

		return nums
	}
}

func hasDuplicates(arr []int) bool {
	mapper := make(map[int]struct{})
	for _, num := range arr {
		if _, ok := mapper[num]; ok {
			return true
		}
		mapper[num] = struct{}{}
	}
	return false
}

func parseNumbers(arr []string) ([]int, error) {
	res := make([]int, len(arr))

	for i, num := range arr {
		parsedNumber, err := strconv.Atoi(num)
		if err != nil {
			return []int{}, fmt.Errorf("not a number")
		}
		res[i] = parsedNumber
	}

	return res, nil
}

func readInput(reader *bufio.Reader) string {
	text, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}

	return strings.TrimRight(text, "\n")
}
