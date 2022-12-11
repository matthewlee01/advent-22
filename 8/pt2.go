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
	i := 0
	grid := make([][]int, 0)
	for err != io.EOF && len(text) > 1 {
		tokens := strings.Split(text, "")
		tokens = tokens[:len(tokens)-1]
		grid = append(grid, make([]int, 0))
		for _, token := range tokens {
			digit, _ := strconv.Atoi(token)
			grid[i] = append(grid[i], digit)
		}
		i++
		text, err = reader.ReadString('\n')
	}

	m := len(grid)
	n := len(grid[0])
	max := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			score := 1
			h := grid[i][j]
			var k int
			for k = 1; k < m-i-1; k++ {
				if grid[i+k][j] >= h {
					break
				}
			}
			score *= k
			for k = 1; k < n-j-1; k++ {
				if grid[i][j+k] >= h {
					break
				}
			}
			score *= k
			for k = 1; k < i; k++ {
				if grid[i-k][j] >= h {
					break
				}
			}
			score *= k
			for k = 1; k < j; k++ {
				if grid[i][j-k] >= h {
					break
				}
			}
			score *= k
			if score > max {
				max = score
			}
		}
	}
	fmt.Println("max:", max)
}
