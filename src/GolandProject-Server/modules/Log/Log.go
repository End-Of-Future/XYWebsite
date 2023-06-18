package Log

import "fmt"

const (
	COLOUR_BLACK   = "\x1B[30m"
	COLOUR_RED     = "\x1B[31m"
	COLOUR_GREEN   = "\x1B[32m"
	COLOUR_YELLOW  = "\x1B[33m"
	COLOUR_BLUE    = "\x1B[34m"
	COLOUR_MAGENTA = "\x1B[35m"
	COLOUR_CYAN    = "\x1B[36m"
	COLOUR_WHITE   = "\x1B[37m"
	RESET          = "\x1B[0m"
)

func Send(a ...any) {
	for i, _ := range a {
		fmt.Print(a[i])
		if i != len(a)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
func Success(value string) {
	fmt.Println(COLOUR_GREEN + value + RESET)
}
func Warn(value string) {
	fmt.Println(COLOUR_YELLOW + "WARN:" + RESET + value)
}
func Error(value string) {
	fmt.Println(COLOUR_RED + "ERR:" + RESET + value)
}
func Fatal(err error) {
	fmt.Println(COLOUR_RED + "Fatal:" + RESET + err.Error())
}
