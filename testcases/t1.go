// -build gomacro_1

package main

func Who(n *float64, i int, w int) int {
	switch w {
	case 0:
		return Tail0(n, i)
	case 1:
		return Tail1(n, i)
	case 2:
		return Tail2(n, i)
	case 3:
		return Tail3(n, i)
	default:
		panic("")
	}
}

func Tail0(n *float64, i int) int {

	(*n)+=(*n)
	(*n)++

	if i <= 0 {
		return 0
	}
	return i + Tail0(n, i-1)
}

func Tail1(n *float64, i int) int {

	(*n)+=(*n)
	(*n)++

	if i > 0 {
		return i + Tail1(n, i-1)
	}
	return 0
}

func Tail2(n *float64, i int) int {

	(*n)+=(*n)
	(*n)++

	if i == 0 {
		return 0
	}

	if (i & 1) == 1 {
		return i + Tail0(n, i-1)
	} else {
		return i + Tail1(n, i-1)
	}
}

func Tail3(n *float64, i int) int {
	return Who(n, i, i % 3)
}

func main() {
	for j := 0; j <= 3; j++ {
	for i := 0; i < 10; i++ {
		f := 0.0
		o := Who(&f, i, j)
		r := 0.0
		q := Who(&r, i, 0)

//		print(f)
//		print("..")
//		print(o)
//		print("\n")

		if f + 0.1 < r || f > 0.1 + r {
			panic("side effect")
		}
		if o != q {
			panic("return val")
		}

	}}
}
