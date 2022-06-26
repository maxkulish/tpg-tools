package main

import (
	"findgo"
	"fmt"
)

func main() {
	fmt.Println(findgo.CountGoFiles("/tmp", 0))
}
