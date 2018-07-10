package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type MyFunc func(a,b int) bool

func HeapMinFunc(a,b int) bool {
	return a > b
}

func HeapMaxFunc(a,b int) bool {
	return a < b
}

func Parent(i int) int {
	if i ==0 {
		return 0
	}
	return (i-1)/2
}

func Left(i int) int {
	return i*2 +1
}

func Right(i int) int {
	return i*2 + 2
}


type Heap struct {
	Arr []int
	F MyFunc
	Length int
}

func CreateHeap(f MyFunc) *Heap {
	h := Heap{}
	h.Arr = make([]int,0)
	h.F = f
	return &h
}


func (h *Heap) Insert(val int) {
	h.Arr = append(h.Arr,val)
	h.Length++
	h.Heapify(h.Length-1)
}

func (h* Heap) RemoveFromHeap() int {
	val1 := h.Arr[0]

	h.Arr[0] = h.Arr[h.Length-1]

	h.Length--

	h.Arr = h.Arr[:h.Length]

	h.ReverseHepify(0)


	return val1
}

func (h *Heap) Heapify(index int) {
	if index == 0 {
		return
	}
	p := Parent(index)
	if h.F(h.Arr[p],h.Arr[index]) {
		tmp := h.Arr[p];h.Arr[p]=h.Arr[index];h.Arr[index] = tmp
		h.Heapify(p)
	}
	return
}

func (h *Heap) ReverseHepify(index int) {
	if index >= h.Length {
		return
	}
	max := -1
	l := Left(index)
	r := Right(index)
	if l >= h.Length {
		return
	}
	if r >= h.Length {
		max = l
		//return
	}else {
		max = r
		if h.F(h.Arr[r],h.Arr[l]) {
			max = l
		}
	}
	if (h.F(h.Arr[index],h.Arr[max])) {
		tmp := h.Arr[max];h.Arr[max]=h.Arr[index];h.Arr[index] = tmp
		h.ReverseHepify(max)
	}
	return
}

func PritMedian(max , min *Heap, item int) {
	if min.Length != 0 && item >= min.Arr[0] {
		min.Insert(item)
		if min.Length - max.Length >= 2 {
			val := min.RemoveFromHeap()
			max.Insert(val)
		}
	}else {
		max.Insert(item)
		if max.Length - min.Length >= 2 {
			val := max.RemoveFromHeap()
			min.Insert(val)
		}
	}

	if max.Length == min.Length {
		fmt.Printf("%.1f\n",float64(float64(max.Arr[0]) + float64(min.Arr[0]))/2)
	}else {
		if min.Length > max.Length {
			fmt.Printf("%.1f\n",float64(min.Arr[0]) )
		}else {
			fmt.Printf("%.1f\n",float64(max.Arr[0]) )
		}

	}
//	fmt.Println(max.Arr)
//	fmt.Println(min.Arr)
}


func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	//var a []int32

	minheap := CreateHeap(HeapMinFunc)
	maxheap := CreateHeap(HeapMaxFunc)

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		aItem := int(aItemTemp)
		PritMedian(maxheap , minheap, aItem);
	}

	return

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
