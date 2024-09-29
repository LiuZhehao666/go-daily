package main

func main() {
	const (
		a = iota
		b
		c
	)
	println(a, b, c)

	const (
		d = iota
		e
		f
	)
	println(d, e, f)
}
