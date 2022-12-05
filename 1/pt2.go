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
	maxSum2 := 0
	maxSum3 := 0
	curSum := 0
	text, err := reader.ReadString('\n')
	for err != io.EOF {
		if len(text) == 1 {
			if curSum > maxSum3 {
				maxSum3 = curSum
			}
			if maxSum3 > maxSum2 {
				temp := maxSum2
				maxSum2 = maxSum3
				maxSum3 = temp
			}
			if maxSum2 > maxSum {
				temp := maxSum
				maxSum = maxSum2
				maxSum2 = temp
			}
			curSum = 0
		} else {
			cur, _ := strconv.Atoi(text[:len(text)-1])
			curSum += cur
		}
		text, err = reader.ReadString('\n')
	}
	fmt.Println("max:", maxSum+maxSum2+maxSum3)
}
