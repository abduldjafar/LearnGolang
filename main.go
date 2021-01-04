package main

import (
	"LearnGolang/interfaces"
	"fmt"
)

func main() {
	data := interfaces.React{}
	data.Width = 2.0
	data.Hight = 5.0

	fmt.Println(data.Area())
}
