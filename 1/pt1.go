package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input:\n")
	maxSum := 0

	curSum := 0
	text, err := reader.ReadString('\n')
	for err != io.EOF {
		if len(text) == 1 {
			if curSum > maxSum {
				maxSum = curSum
			}
			curSum = 0
		} else {
			cur, _ := strconv.Atoi(text[:len(text)-1])
			curSum += cur
		}
		text, err = reader.ReadString('\n')
	}
	fmt.Println("max:", maxSum)
}
