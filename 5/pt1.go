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
	var stacks = make([][]byte, len(text)/4)
	for text[1] != '1' {
		for i := 1; i < len(text); i += 4 {
			if text[i] != ' ' {
				stacks[(i-1)/4] = append([]byte{text[i]}, stacks[(i-1)/4]...)
			}
		}
		text, err = reader.ReadString('\n')
	}
	reader.ReadString('\n')
	text, err = reader.ReadString('\n')
	for err != io.EOF && len(text) > 1 {
		text = text[5 : len(text)-1]
		splitFrom := strings.Split(text, " from ")
		splitTo := strings.Split(splitFrom[1], " to ")
		quantity, _ := strconv.Atoi(splitFrom[0])
		pos_source, _ := strconv.Atoi(splitTo[0])
		pos_dest, _ := strconv.Atoi(splitTo[1])
		ptr_source := &stacks[pos_source-1]
		ptr_dest := &stacks[pos_dest-1]
		for i := 0; i < quantity; i++ {
			source := *ptr_source
			dest := *ptr_dest
			c := source[len(source)-1]
			*ptr_source = source[:len(source)-1]
			*ptr_dest = append(dest, c)
		}
		text, err = reader.ReadString('\n')
	}

	fmt.Println("final config:")
	fmt.Println(stacks)
	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
}
