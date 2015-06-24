package main

import (
	"fmt"
	"github.com/hnw/go-smc/smc"
)

func main() {
	smc.Open()
	fmt.Printf("%0.1f℃\n", smc.GetTemperature(smc.KEY_CPU_TEMP))
	smc.Close()
}
