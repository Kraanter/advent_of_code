package days

type Day interface {
	GetDayNumber() int
	First([]string) int
	Second([]string) int
}
