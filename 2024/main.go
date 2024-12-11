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
	days.Day6,
	days.Day7,
	days.Day8,
	days.Day9,
	days.Day10,
}

func main() {
	realPtr := flag.Bool("real", false, "Flag to use the actual data from AoC instead of the test data")
	dayPtr := flag.Int("day", -1, "Day number to run")
	flag.Parse()
	util.IsTestInput = !*realPtr
	if *dayPtr != -1 {
		runDay(solutions[*dayPtr-1])
	} else {
		for _, day := range solutions {
			runDay(day)
		}
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
