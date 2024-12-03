package days

type DaySolution interface {
	First([]string) int
	Second([]string) int
}

type Day struct {
	DaySolution
	DayNr int
}
