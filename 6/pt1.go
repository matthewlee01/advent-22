package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input:\n")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	for i := 4; i < len(text); i++ {
		buf := text[i-4 : i]
		unique := true
		fmt.Println(buf)
	out:
		for j, c1 := range buf[:3] {
			for _, c2 := range buf[j+1:] {
				if c1 == c2 {
					fmt.Println(string(c1))
					unique = false
					i += j
					break out
				}
			}
		}
		if unique {
			fmt.Println(i)
			break
		}
	}
}
