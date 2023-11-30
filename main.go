package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/tribals/hexlet-go/greeting"
)

func main() {
	fmt.Println(color.Red(greeting.Hello()))
}
