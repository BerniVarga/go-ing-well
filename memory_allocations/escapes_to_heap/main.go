package main

func main() {
	s := 45000
	println("salary", s, "with address", &s)

	bonus := calculateBonus(s)
	println("bonus", bonus, "with address", &bonus)
}

//go:noinline
func calculateBonus(s int) *int {

	multiplier := 80 // 80%
	bonus := s * multiplier / 100

	// removed printing if you'd like to keep your program more simple when building with gcflags
	println("	salary", s, "with address", &s)
	println("	multiplier", multiplier, "%", "with address", &multiplier)
	println("	bonus", bonus, "with address", &bonus)

	return &bonus
}
