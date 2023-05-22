package main

// Interesting examples with string

// 1. No allocation with byte appends:
func main() {
	s := "have fun"
	s2 := atDeliveroo([]byte(s))

	println(string(s2))
}

//go:noinline
func atDeliveroo(n []byte) []byte {
	n2 := append(n, []byte(" at Deliveroo")...)
	return n2
}

/*
//2. No allocation with string join (that uses the builder below):
func main() {
	s := "have fun"
	s2 := atDeliveroo(s)
	println(s2)
}

//go:noinline
func atDeliveroo(n string) string{
	s := strings.Join([]string{n, "at Deliveroo"}, " ")
	return s
}
*/

/*
//3. Allocation is happening with default string concatenation or fmt.Println():
func main() {
	s := "have fun"
	s2 := atDeliveroo(s)

	println(s2)
}

//go:noinline
func atDeliveroo(s string) string {
	return s + " at Deliveroo"
}
*/
