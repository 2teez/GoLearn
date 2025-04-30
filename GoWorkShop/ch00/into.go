package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func random() {
	for range 10 {
		r := rand.Intn(10) + 1
		fmt.Println(r, strings.Repeat("*", r))
	}
}
