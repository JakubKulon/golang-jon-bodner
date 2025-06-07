package main

import (
	"fmt"
	"math"
)

func main() {
	first()
}

func first() {
	var i int = 20
	var f float32 = i

	fmt.Println(i)
	fmt.Println(f)
}

func second() {
	value := 20
	var i int = value
	var r float32 = value
}

func third() {
	var b byte = math.MaxUint8 + 1
	var smallI int32 = math.MaxInt32 + 1
	var bigI uint64 = uint64 + 1
}
