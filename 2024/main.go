package main

import (
	"adc_2024/days"
	"adc_2024/util"
	"embed"
	"fmt"
)

//go:embed input/*
var inputFS embed.FS

var solutions []days.Day = []days.Day{
	days.Day3,
}

func main() {
	util.IsTestInput = true
	for _, day := range solutions {
		runDay(day)
	}
}

func runDay(day days.Day) {
	lines := getDayLines(day.DayNr)
	fmt.Printf("DAY %v | First: %v \n", day.DayNr, day.First(lines))
	fmt.Printf("DAY %v | Second: %v \n", day.DayNr, day.Second(lines))
}

func getDayLines(day int) []string {
	input := util.ReadDayInput(day, inputFS)

	return util.BytesToLines(input)
}
