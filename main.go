package main

import (
    "flag" // 追加する
    "fmt"
)

func main() {
    flag.Parse()
    arg := flag.Arg(0)
    fmt.Printf("Hello %s\n", arg)
}
