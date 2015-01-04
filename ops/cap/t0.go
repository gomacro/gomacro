package main

func macro(s []) int {
	return cap(s)
}

func main() {
	a := []int{1,3,3,7}
	b := []byte("1337")
	c := []bool{true, true, true, true, true, false}

	print(macro(a))
	print(macro(b))
	print(macro(c))

}
