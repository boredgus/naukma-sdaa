package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sdaa/lab03/src"
	"strconv"
	"strings"
)

func main() {
	removeOldTreeSnapshots()

	reader := bufio.NewReader(os.Stdin)
	n, m := scanNAndM(reader)

	firstTree := &src.RBTree{}
	fmt.Printf("\nFill first tree with values:\n\n")
	fillTreeWithNElements(
		reader,
		firstTree,
		n,
		func(addedElement *src.RBTreeNode) {
			current := addedElement
			for {
				if current.Parent == nil {
					break
				}
				current = current.Parent
			}
			for {
				if current.Right == nil {
					break
				}
				current = current.Right
			}
			fmt.Printf("max element of tree after inserting: %s\n", current)
		},
		"first",
	)

	fmt.Printf("\nFill second tree with values:\n\n")
	secondTree := &src.RBTree{}
	fillTreeWithNElements(
		reader,
		secondTree,
		m,
		func(addedNode *src.RBTreeNode) {
			fmt.Printf("parent node of added node: %s\n", addedNode.Parent)
		},
		"second",
	)
}

func scanNAndM(reader *bufio.Reader) (n int, m int) {
	for {
		fmt.Println(`
Enter N and M separated by space, where:
	N - amount of elements to add into first tree (N > 2)
	M - amount of elements to add into second tree (M > 2)
Example: "6 8"`)

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
		if n < 3 || n > 255 {
			fmt.Println("> Invalid format of data: N should be in range [3,256)")
			continue
		}

		m, err = strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("> Invalid format of data: M is not a number. Try again.")
			continue
		}
		if m < 3 || m > 255 {
			fmt.Println("> Invalid format of data: M should be in range [3,256)")
			continue
		}

		break
	}

	return n, m
}

func fillTreeWithNElements(
	reader *bufio.Reader,
	tree *src.RBTree,
	desiredCountOfElements int,
	actionOnAdd func(addedNode *src.RBTreeNode),
	treeName string,
) {
	countOfAddedElements := 0

	for countOfAddedElements != desiredCountOfElements {
		fmt.Printf(" (%v of %v) Enter value (integer) of new element:", countOfAddedElements+1, desiredCountOfElements)
		inputLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("> Failed to scan your input. Try again.")
			continue
		}

		inputLine = strings.TrimSpace(inputLine)
		value, err := strconv.Atoi(inputLine)
		if err != nil {
			fmt.Println("> Invalid format of data: N is not an integer. Try again.")
			continue
		}

		addedElement, err := tree.Insert(value)
		if err != nil {
			fmt.Printf("> Failed to add element to tree: %v. Try again.\n", err)
			continue
		}

		actionOnAdd(addedElement)
		countOfAddedElements++

		src.DrawGraph(tree, treeName, countOfAddedElements)
	}
}

func removeOldTreeSnapshots() {
	files, err := filepath.Glob("lab3_tree_*")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}

}
