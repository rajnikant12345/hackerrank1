package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Count int
	N []Node
}

var Parent Node



func AddToTrie(str string) {
	tmp := &Parent
	for _,i:= range str {
		tmp.N[i-'a'].Count++
		if len(tmp.N[i-'a'].N) == 0 {
			tmp.N[i-'a'].N = make([]Node,26)
		}
		tmp = &tmp.N[i-'a']
	}
	return
}

func FindInTrie(str string) int {
	count := 0
	tmp := &Parent
	for _,i:= range str {
		if tmp.N[i-'a'].Count == 0 {
			return 0
		}else {
			count = tmp.N[i-'a'].Count
			tmp = &tmp.N[i-'a']
		}
	}
	return count
}


func main() {

	Parent.N = make([]Node,26)

	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	for nItr := 0; nItr < int(n); nItr++ {
		opContact := strings.Split(readLine(reader), " ")

		op := opContact[0]

		contact := opContact[1]
		if op == "add" {
			AddToTrie(contact)
		}else {
			fmt.Println(FindInTrie(contact))
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
