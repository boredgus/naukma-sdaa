package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sdaa/common"
	"sdaa/lab04/src"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var minimalNumber *int
	minimalNumber = new(int)
	*minimalNumber = 3
	n := common.ScanNumber(reader, "Enter a number, which represents amount of points to enter: ", minimalNumber)
	points := scanPoints(reader, n, 100)

	for idx, point := range points {
		if idx > 1 {
			direction := src.GetTurnDirection(points[idx-2], points[idx-1], points[idx])
			fmt.Println(direction)
		}
		fmt.Println(point.String())
	}

	src.DrawPoints(points)
}

func scanPoints(reader *bufio.Reader, desiredAmount int, maxValue float64) []src.Point {
	points := make([]src.Point, 0, desiredAmount)

	for len(points) < desiredAmount {
		fmt.Printf(
			" (%v of %v) Enter 2 float values separated by space (abs(value) < %.2f): ",
			len(points)+1, desiredAmount, maxValue,
		)

		inputLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("> Failed to scan your input. Try again.")
			continue
		}

		inputLine = strings.TrimSpace(inputLine)
		arguments := strings.Split(inputLine, " ")
		if len(arguments) != 2 {
			fmt.Println("> Invalid format of data. There should be exactly 2 arguments. Try again.")
			continue
		}

		x, err := strconv.ParseFloat(arguments[0], 64)
		if err != nil {
			fmt.Println("> Invalid format of data: first argument is not a number with floating point. Try again.")
			continue
		}
		if math.Abs(x) > maxValue {
			fmt.Printf(
				"> Invalid format of data: first argument is out of range [-%s,%s]. Try again.\n",
				maxValue, maxValue,
			)
			continue
		}

		y, err := strconv.ParseFloat(arguments[1], 64)
		if err != nil {
			fmt.Println("> Invalid format of data: second argument is not a number with floating point. Try again.")
			continue
		}
		if math.Abs(y) > maxValue {
			fmt.Printf(
				"> Invalid format of data: second argument is out of range [-%s,%s]. Try again.\n",
				maxValue, maxValue,
			)
			continue
		}

		points = append(points, src.NewPoint(x, y))
	}

	return points
}
