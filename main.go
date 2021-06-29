package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	c, _ := cpu.Info()

	fmt.Println("Hello, world.")
	fmt.Println(c)
}
