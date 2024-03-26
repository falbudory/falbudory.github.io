package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

var colors = []string{Red, Green, Yellow, Blue, Purple, Cyan, White}

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		color := colors[rand.Intn(len(colors))]
		text := "          ❤ Anh yêu em bé của anh rất nhiềuuu ❤ "
		fmt.Printf("%s%s%s\n", color, text, Reset)
		time.Sleep(200 * time.Millisecond)
	}
}
