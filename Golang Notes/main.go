package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)
	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)


}