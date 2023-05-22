package main

func main() {
	s := 45000
	println("salary", s, "with address", &s)
	calculateBonus(&s)
	println("bonus", s, "with address", &s)
}

//go:noinline
func calculateBonus(s *int) {

	multiplier := 80 // 80%
	*s = (*s) * multiplier / 100

	// removed printing if you'd like to keep your program more simple when building with gcflags
	println("	salary", s, "with address", &s)
	println("	multiplier", multiplier, "%", "with address", &multiplier)
	println("	bonus", s, "with address", &s)
}
