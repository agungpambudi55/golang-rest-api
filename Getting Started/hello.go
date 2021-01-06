package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println("Today's Quote")
    fmt.Println(quote.Go())
}