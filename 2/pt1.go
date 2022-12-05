package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	symbolValue := make(map[byte]int)
	symbolValue['A'] = 1
	symbolValue['B'] = 2
	symbolValue['C'] = 3
	symbolValue['X'] = 1
	symbolValue['Y'] = 2
	symbolValue['Z'] = 3
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input:\n")
	totalScore := 0
	text, err := reader.ReadString('\n')
	for len(text) > 2 && err != io.EOF {
		s1 := text[0]
		s2 := text[2]

		if (s1 == 'A' && s2 == 'Y') ||
			(s1 == 'B' && s2 == 'Z') ||
			(s1 == 'C' && s2 == 'X') {
			totalScore += 6
		}
		if symbolValue[s1] == symbolValue[s2] {
			totalScore += 3
		}
		totalScore += symbolValue[s2]
		text, err = reader.ReadString('\n')
	}
	fmt.Println("total score:", totalScore)
}
