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
	hx, hy, tx, ty := 0, 0, 0, 0
	visited := make(map[string]bool)
	for err != io.EOF && len(text) > 1 {
		instr := strings.Split(text, " ")
		dir := instr[0]
		n, _ := strconv.Atoi(instr[1][:len(instr[1])-1])

		for i := 0; i < n; i++ {
			switch dir {
			case "U":
				hy += 1
			case "R":
				hx += 1
			case "D":
				hy -= 1
			case "L":
				hx -= 1
			}
			dist := int(math.Abs(float64(hy-ty)) + math.Abs(float64(hx-tx)))
			threshold := 1
			if dist > 2 {
				threshold = 0
			}
			if hy-ty > threshold {
				ty += 1
			}
			if hy-ty < -threshold {
				ty -= 1
			}
			if hx-tx > threshold {
				tx += 1
			}
			if hx-tx < -threshold {
				tx -= 1
			}
			fmt.Println(hx, hy, tx, ty)
			visited[strconv.Itoa(tx)+","+strconv.Itoa(ty)] = true
		}
		text, err = reader.ReadString('\n')
	}
	fmt.Println(visited)
	fmt.Println(len(visited))
}
