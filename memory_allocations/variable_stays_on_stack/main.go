package main

func main() {
	salary := 45000
	println("salary", salary, "at address", &salary)

	bonus := calculateBonus(salary)
	println("bonus", bonus, "at address", &bonus)

	bonus = addALittleExtra(bonus)
	println("bonus", bonus, "at address", &bonus)
}

//go:noinline
func calculateBonus(s int) int {
	multiplier := 80 // 80%
	bonus := s * multiplier / 100

	// removed printing if you'd like to keep your program more simple when building with gcflags
	println("	salary", s, "with address", &s)
	println("	multiplier", multiplier, "%", "with address", &multiplier)
	println("	bonus", bonus, "with address", &bonus)

	return bonus
}

//go:noinline
func addALittleExtra(bonus int) int {
	bonus = bonus * 120 / 100
	println("	bonus", bonus, "with address", &bonus)
	return bonus
}
