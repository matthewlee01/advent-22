package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input:\n")
	text, err := reader.ReadString('\n')
	for err != io.EOF {
		text, err = reader.ReadString('\n')
	}
	fmt.Println("max:", text)
}
