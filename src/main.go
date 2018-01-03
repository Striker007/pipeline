package main

import (
    "fmt"
    "chapter1"
    )

func main() {
    fmt.Println("start")
    ten := chapter1.New_Dollar(5)

    result := ten.Times(2)
    fmt.Println(result)
}

