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
	m := make(map[byte]int)
	line := 0
	for err != io.EOF {
		pos := line % 3
		for i := 0; i < len(text); i++ {
			if count, ok := m[text[i]]; ok {
				if count == 1 && pos == 1 {
					m[text[i]]++
				} else if count == 2 && pos == 2 {
					score += strings.Index(alphabet, string(text[i]))
					m[text[i]]++
					break
				}
			} else {
				if pos == 0 {
					m[text[i]] = 1
				}
			}
		}
		if pos == 2 {
			m = make(map[byte]int)
		}
		line++
		text, err = reader.ReadString('\n')
	}
	fmt.Println("score:", score)
}
