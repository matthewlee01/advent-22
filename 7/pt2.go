package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const TOTAL_SPACE int = 70000000
const REQUIRED_SPACE int = 30000000

type Dir struct {
	filesize int
	children map[string]*Dir
}

func traverseDir(dir Dir, depth int) (sizes []int) {
	size := dir.filesize
	child_sizes := []int{}
	for name, child := range dir.children {
		if name != ".." {
			fmt.Println(strings.Repeat(" ", depth), name)
			child_sizes = append(traverseDir(*child, depth+1), child_sizes...)
			size += child_sizes[0]
		}
	}
	fmt.Println(strings.Repeat(" ", depth), "total filesize:", size)
	return append([]int{size}, child_sizes...)
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
	sizes := traverseDir(root, 0)
	free := TOTAL_SPACE - sizes[0]
	min_to_delete := REQUIRED_SPACE - free
	min_size := TOTAL_SPACE
	for _, size := range sizes {
		if size > min_to_delete && size < min_size {
			min_size = size
		}
	}
	fmt.Println(min_size)
}
