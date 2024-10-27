package main

import (
	"bufio"
	"fmt"
	"os"
	"sdaa/lab02/src"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := scanNAndM(reader)

	hashTable := src.NewHashTable(256)
	fillHashTable(reader, hashTable, n)

	hashTable.PrintSelf()

	keysToFind := scanKeysToFind(reader, m)

	fmt.Println("\nResults of search:")
	for _, key := range keysToFind {
		value, found := hashTable.Search(key)

		if found {
			fmt.Printf("Key: %v\t Value: '%v'\n", key, value)
		} else {
			fmt.Printf("There is no key '%v' in hash table\n", key)
		}
	}
}

func scanNAndM(reader *bufio.Reader) (n int, m int) {
	for {
		fmt.Println(`
Enter N and M separated by space, where:
	N - amount of pairs to add to hash
	M - amount of keys to search in hash table
Example: "4 2"`)

		inputLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("> Failed to scan your input. Try again.")
			continue
		}

		inputLine = strings.TrimSpace(inputLine)
		parts := strings.Split(inputLine, " ")
		if len(parts) != 2 {
			fmt.Println("> Invalid format of data. There should be exactly 2 arguments. Try again.")
			continue
		}

		n, err = strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("> Invalid format of data: N is not a number. Try again.")
			continue
		}
		if n < 2 || n > 255 {
			fmt.Println("> Invalid format of data: N should be in range [2,256)")
			continue
		}

		m, err = strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("> Invalid format of data: M is not a number. Try again.")
			continue
		}
		if m < 2 || m > 255 {
			fmt.Println("> Invalid format of data: M should be in range [2,256)")
			continue
		}

		break
	}

	return n, m
}

func fillHashTable(reader *bufio.Reader, hashTable *src.HashTable, desiredCount int) {
	countOfAddedPairs := 0

	fmt.Printf("\nEnter %v pairs consisting of number and string in format '1.1 value'.\n", desiredCount)
	for countOfAddedPairs < desiredCount {
		fmt.Printf("> (%v from %v):\t", countOfAddedPairs+1, desiredCount)

		inputLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to scan your input. Try again.")
			continue
		}

		inputLine = strings.TrimSpace(inputLine)
		parts := strings.SplitN(inputLine, " ", 2)
		if len(parts) != 2 {
			fmt.Println("> Invalid format of data. There should be exactly 2 arguments. Try again.")
			continue
		}

		key, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid format of data: key is not a number with float point. Try again.")
			continue
		}

		hashTable.Put(key, parts[1])
		countOfAddedPairs++
	}
}

func scanKeysToFind(reader *bufio.Reader, desiredCount int) []float64 {
	keysToFind := make([]float64, 0, desiredCount)
	fmt.Printf("\nEnter %v keys (number with floating point) to find in hash table.\n", desiredCount)
	for len(keysToFind) < desiredCount {
		fmt.Printf("> (%v from %v):\t", len(keysToFind)+1, desiredCount)

		inputLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to scan your input. Try again.")

			continue
		}

		inputLine = strings.TrimSpace(inputLine)
		key, err := strconv.ParseFloat(inputLine, 64)
		if err != nil {
			fmt.Println("Invalid format of data: key is not a number with float point. Try again.")

			continue
		}

		keysToFind = append(keysToFind, key)
	}

	return keysToFind
}
