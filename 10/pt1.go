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
	var err error
	text := "  "
	total := 0
	i := 1
	x := 1
	v := 0
	wait := 0
	for err != io.EOF && len(text) > 1 {
		if (i+20)%40 == 0 && i <= 240 {
			total += i * x
		}
		if wait == 0 {
			text, err = reader.ReadString('\n')

			if text[0] == 'a' {
				vstring := strings.Split(text, " ")[1]
				vstring = vstring[:len(vstring)-1]
				v, _ = strconv.Atoi(vstring)
				wait = 1
			}
		} else {
			wait--
			if v != 0 && wait == 0 {
				x += v
				v = 0
			}
		}
		i++
	}
	fmt.Println("total:", total)
}
