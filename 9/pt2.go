package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input:\n")
	text, err := reader.ReadString('\n')
	xs := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ys := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited := make(map[string]bool)
	for err != io.EOF && len(text) > 1 {
		instr := strings.Split(text, " ")
		dir := instr[0]
		n, _ := strconv.Atoi(instr[1][:len(instr[1])-1])

		for i := 0; i < n; i++ {
			switch dir {
			case "U":
				ys[0] += 1
			case "R":
				xs[0] += 1
			case "D":
				ys[0] -= 1
			case "L":
				xs[0] -= 1
			}
			for j := 1; j < 10; j++ {
				// fmt.Print(j-1, xs[j-1], ys[j-1], " | ")
				dist := int(math.Abs(float64(ys[j-1]-ys[j])) + math.Abs(float64(xs[j-1]-xs[j])))
				threshold := 1
				if dist > 2 {
					threshold = 0
				}
				if ys[j-1]-ys[j] > threshold {
					ys[j] += 1
				}
				if ys[j-1]-ys[j] < -threshold {
					ys[j] -= 1
				}
				if xs[j-1]-xs[j] > threshold {
					xs[j] += 1
				}
				if xs[j-1]-xs[j] < -threshold {
					xs[j] -= 1
				}
			}
			// fmt.Println("9", xs[9], ys[9])
			visited[strconv.Itoa(xs[9])+","+strconv.Itoa(ys[9])] = true
		}
		text, err = reader.ReadString('\n')
	}
	fmt.Println(len(visited))
}
