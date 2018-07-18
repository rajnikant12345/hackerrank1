package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var count int64

func Merge(arr1 []int32, arr2 []int32) []int32 {
	if arr1 == nil && arr2 == nil {
		return nil
	}
	if arr1 == nil {
		return arr2
	}
	if arr2 == nil {
		return arr1
	}
	tmp := make([]int32,len(arr1)+len(arr2))
	i := 0
	j := 0
	for k,_ := range tmp {
		if i >= len(arr1) {
			tmp[k] = arr2[j]
			j++
		}else if j >= len(arr2) {
			tmp[k] = arr1[i]
			i++
		}else if arr1[i] > arr2[j] {
			tmp[k] = arr2[j]
			//fmt.Println(i ,j ,k)
			count += int64(j+len(arr1)-k)
			j++

		}else {
			tmp[k] = arr1[i]
			i++
		}
	}

	return tmp

}

func MergeSort(arr []int32) []int32 {

	if len(arr) <= 1 {
		return arr
	}
	middle := (len(arr)-1)/2
	arr1 := MergeSort(arr[0:middle+1])
	arr2 := MergeSort(arr[middle+1:len(arr)])
	return Merge(arr1, arr2)
}

// Complete the countInversions function below.
func countInversions(arr []int32) int64 {
	arr = MergeSort(arr)
	//fmt.Println(arr)
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	//stdout := os.Stdout

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := countInversions(arr)

		fmt.Fprintf(writer, "%d\n", result)
		count = 0
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
