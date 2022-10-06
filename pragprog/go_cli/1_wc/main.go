package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)


func main() {
    lines := flag.Bool("l", false, "Count Lines")
    flag.Parse()

    num := count(os.Stdin, *lines)
    fmt.Println(num)
}

func count(r io.Reader, countLines bool) int {
    scanner := bufio.NewScanner(r)

    if !countLines {
        scanner.Split(bufio.ScanWords)
    }

    wc := 0
    for scanner.Scan() {
        wc++
    }

    return wc
}

