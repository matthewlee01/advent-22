package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input:\n")
	text, err := reader.ReadString('\n')
	const alphabet string = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	score := 0
	for err != io.EOF {
		m := make(map[byte]bool)
		for i := 0; i < len(text)/2; i++ {
			m[text[i]] = true
		}
		for i := len(text) / 2; i < len(text); i++ {
			if _, ok := m[text[i]]; ok {
				score += strings.Index(alphabet, string(text[i]))
				break
			}
		}
		text, err = reader.ReadString('\n')
	}
	fmt.Println("score:", score)
}
