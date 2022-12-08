package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input:\n")
	text, err := reader.ReadString('\n')
	total := 0
	for err != io.EOF && len(text) > 1 {
		pair := strings.Split(text[:len(text)-1], ",")
		rangeA := strings.Split(pair[0], "-")
		rangeB := strings.Split(pair[1], "-")
		startA, _ := strconv.Atoi(rangeA[0])
		endA, _ := strconv.Atoi(rangeA[1])
		startB, _ := strconv.Atoi(rangeB[0])
		endB, _ := strconv.Atoi(rangeB[1])
		if (startA <= startB && endA >= endB) ||
			(startB <= startA && endB >= endA) {
			total++
		}
		text, err = reader.ReadString('\n')
	}
	fmt.Println("total:", total)
}
