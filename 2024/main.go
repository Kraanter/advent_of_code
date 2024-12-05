package main

import (
	"adc_2024/days"
	"adc_2024/util"
	"embed"
	"flag"
	"fmt"
)

//go:embed input/*
var inputFS embed.FS

var solutions []days.Day = []days.Day{
	days.Day1,
	days.Day2,
	days.Day3,
	days.Day4,
	days.Day5,
}

func main() {
	realPtr := flag.Bool("real", false, "Flag to use the actual data from AoC instead of the test data")
	flag.Parse()
	util.IsTestInput = !*realPtr
	for _, day := range solutions {
		runDay(day)
	}
}

func runDay(day days.Day) {
	fmt.Println("------------------")
	dayNumber := day.GetDayNumber()
	lines := getDayLines(dayNumber)
	fmt.Printf("DAY %v | First: %v \n", dayNumber, day.First(lines))
	fmt.Printf("DAY %v | Second: %v \n", dayNumber, day.Second(lines))
	fmt.Println("------------------")
}

func getDayLines(day int) []string {
	input := util.ReadDayInput(day, inputFS)

	return util.BytesToLines(input)
}
