package main

import (
	"bufio"
	"fmt"
	"os"
	"sdaa/common"
	"sdaa/lab01/src"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	minAmount := 2
	amount := common.ScanNumber(reader, "Enter amount of input numbers: ", &minAmount)
	numbers := common.ScanUniqueNumbers(reader, amount)

	list := src.DoublyLinkedList{}

	for _, num := range numbers {
		list.AddNode(src.NewNode(num))
	}

	fmt.Println("sum: ", list.Sum())
	fmt.Println("average: ", list.Average())
	fmt.Println("smallest: ", list.GetSmallest(3))
	fmt.Println("largest: ", list.GetLargest(3))
	fmt.Println("(n/2)th element: ", list.GetNthElement(amount/2))
}
