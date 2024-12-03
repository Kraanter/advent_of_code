package main

import (
	"adc_2024/day3"
	"adc_2024/util"
	"embed"
	"fmt"
)

//go:embed input/*
var input embed.FS

func main() {
	input, err := input.ReadFile("input/day3")
	if err != nil {
		println("error: " + err.Error())
		return
	}

	lines := util.BytesToLines(input)
	fmt.Printf("DAY 3 | First: %v \n", day3.First(lines))
	fmt.Printf("DAY 3 | Second: %v \n", day3.Second(lines))
}
