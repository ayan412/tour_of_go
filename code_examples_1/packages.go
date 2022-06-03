package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(100000))
	rand.Seed(time.Now().UTC().UnixNano()) // как это должно работать?
}
