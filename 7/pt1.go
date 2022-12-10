package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	filesize int
	children map[string]*Dir
}

func traverseDir(dir Dir, depth int) (size int, sum int) {
	sum = 0
	size = dir.filesize
	for name, child := range dir.children {
		if name != ".." {
			fmt.Println(strings.Repeat(" ", depth), name)
			child_size, child_sum := traverseDir(*child, depth+1)
			size += child_size
			sum += child_sum
		}
	}
	if size <= 100000 {
		sum += size
	}
	fmt.Println(strings.Repeat(" ", depth), "total filesize:", size, sum)
	return size, sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input:\n")
	text, err := reader.ReadString('\n')
	root := Dir{children: make(map[string]*Dir)}
	dir := &root
	for err != io.EOF && len(text) > 1 {
		if text[:4] == "$ cd" {
			name := text[5 : len(text)-1]
			if _, ok := (*dir).children[name]; !ok {
				fmt.Println("creating dir", name)
				newdir := Dir{children: make(map[string]*Dir)}
				newdir.children[".."] = dir
				(*dir).children[name] = &newdir
			}
			newdir := (*dir).children[name]
			dir = newdir
			text, err = reader.ReadString('\n')
		} else if text[:4] == "$ ls" {
			text, err = reader.ReadString('\n')
			for text[0] != '$' && len(text) > 1 {
				line := strings.Split(text, " ")
				if line[0] == "dir" {
					name := line[1][:len(line[1])-1]
					newdir := Dir{children: make(map[string]*Dir)}
					newdir.children[".."] = dir
					(*dir).children[name] = &newdir
				} else {
					size, _ := strconv.Atoi(line[0])
					(*dir).filesize += size
				}
				text, err = reader.ReadString('\n')
			}
		}
	}
	fmt.Println("root:")
	_, sum := traverseDir(root, 0)
	fmt.Println("sum:", sum)
}
