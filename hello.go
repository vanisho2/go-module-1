package hello

import (
	"rsc.io/quote/v3"
	"fmt"
)

func Hello() string {
    return quote.HelloV3()
}

func Proverb() string {
    return quote.Concurrency()
}

func SayHello() {
	fmt.Println("Hello world")
}