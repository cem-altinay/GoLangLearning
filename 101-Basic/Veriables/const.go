package main

const sabit1 = "deger 1"

const (
	sabit2 = "deger2"
	sabit3 = "deger3"
)

const (
	sabit5 = iota
	sabit6
	sabit7
)

func main() {
	println(sabit1, sabit2, sabit5, sabit7)
}
