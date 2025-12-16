package calc

func Add(x int, y int) int {
	return x + y
}
func Sub(x int, y int) int {
	return x - y
}
func noThing(x int, y int) int {
	return 0
}
func Calc(do string) func(x int, y int) int {
	switch do {
	case "add":
		return Add
	case "sub":
		return Sub
	default:
		return noThing
	}

}
