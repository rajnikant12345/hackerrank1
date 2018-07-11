package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var F []int32

// Complete the stepPerms function below.
func stepPerms(n int32) int32 {
    if n < 0 {
		return 0
	}
	if n == 0 {
        return 1
	}
    if F[n] != 0 {
        return F[n]
    }
    F[n] = stepPerms(n-1) + stepPerms(n-2) + stepPerms(n-3)
    return F[n]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	s := int32(sTemp)

	for sItr := 0; sItr < int(s); sItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

        F = make([]int32,n+1)
        
		res := stepPerms(n)

		fmt.Fprintf(writer, "%d\n", res)
		
	}

	writer.Flush()
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
