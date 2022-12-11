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
	visible := make([][]bool, m)
	for k := range visible {
		visible[k] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		max := -1
		for j := 0; j < n; j++ {
			if grid[i][j] > max {
				if !visible[i][j] {
					visible[i][j] = true
					total++
				}
				max = grid[i][j]
			}
		}
	}
	for i := 0; i < m; i++ {
		max := -1
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] > max {
				if !visible[i][j] {
					visible[i][j] = true
					total++
				}
				max = grid[i][j]
			}
		}
	}
	for j := 0; j < n; j++ {
		max := -1
		for i := 0; i < m; i++ {
			if grid[i][j] > max {
				if !visible[i][j] {
					visible[i][j] = true
					total++
				}
				max = grid[i][j]
			}
		}
	}
	for j := 0; j < n; j++ {
		max := -1
		for i := m - 1; i >= 0; i-- {
			if grid[i][j] > max {
				if !visible[i][j] {
					visible[i][j] = true
					total++
				}
				max = grid[i][j]
			}
		}
	}
	fmt.Println("total:", total)
}
