package main

import (
	"fmt"

	"github.com/R167/go-remind"
)

func main() {
	_, err := remind.Parse("in 5 minutes")
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello, world!")
}
