package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

var F map[int32]int32
var V map[int32]int32
// Complete the ways function below.
func ways(n int32, coins []int32, j int) int32 {
    if n < 0 {
        return 0
    }
    if n == 0 {
        return 1
    }
    
   /* if F[n] != 0 {
        return F[n]
    }*/
    
    val := int32(0)
    for i,_ := range coins {
        if j > i {
            continue
        }
        val += ways(n-coins[i],coins, i)  
      //  F[n] = val
    }
    return val
}

func main() {
    F = make(map[int32]int32)
    V = make(map[int32]int32)
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    coinsTemp := strings.Split(readLine(reader), " ")

    var coins []int32

    for i := 0; i < int(m); i++ {
        coinsItemTemp, err := strconv.ParseInt(coinsTemp[i], 10, 64)
        checkError(err)
        coinsItem := int32(coinsItemTemp)
        coins = append(coins, coinsItem)
    }

    res := ways(n, coins, 0)

    fmt.Fprintf(writer, "%d\n", res)

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
