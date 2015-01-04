// -build gomacro_1

package main

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

func main() {
// macro vars

	var marg000_i, mret_000 int
	var marg000_n *float64

	var marg001_i, mret_001 int
	var marg001_n *float64

	var marg002_i, mret_002 int
	var marg002_n *float64

	var marg003_i, mret_003 int
	var marg003_n *float64

	var marg5d3_i, marg_5d3_w, mret_5d3 int
	var marg5d3_n *float64
// macro func
	macro := func (z60adfg465as6dg156as165d1g56sa1dg6a1d int) {
		switch z60adfg465as6dg156as165d1g56sa1dg6a1d {
		case 0:
		{
	(*marg000_n)+=(*marg000_n)
	(*marg000_n)++

	if marg000_i <= 0 {
		mret_000 = 0
		return
	}
	mret_000 = marg000_i + Tail0(marg000_n, marg000_i-1)
	return
		}
		case 1:
		{
	(*marg001_n)+=(*marg001_n)
	(*marg001_n)++

	if marg001_i > 0 {
		mret_001 = marg001_i + Tail1(marg001_n, marg001_i-1)
		return
	}
	mret_001 = 0
	return
		}
		case 2:
		{
	(*marg002_n)+=(*marg002_n)
	(*marg002_n)++

	if marg002_i == 0 {
		mret_002 = 0
		return
	}

	if (marg002_i & 1) == 1 {
		mret_002 =  marg002_i + Tail0(marg002_n, marg002_i-1)
	} else {
		mret_002 =  marg002_i + Tail1(marg002_n, marg002_i-1)
	}
	return
		}
		case 3:
		{
	mret_003 = Who(marg003_n, marg003_i, marg003_i % 3)
	return
		}
		case 4:
		{
		switch marg_5d3_w {
		case 0:
			mret_5d3 = Tail0(marg5d3_n, marg5d3_i)
			return
		case 1:
			mret_5d3 = Tail1(marg5d3_n, marg5d3_i)
			return
		case 2:
			mret_5d3 = Tail2(marg5d3_n, marg5d3_i)
			return
		case 3:
			mret_5d3 = Tail3(marg5d3_n, marg5d3_i)
			return
		default:
			panic("")
		}
		}
		}
	}

	for j := 0; j <= 3; j++ {
	for i := 0; i < 10; i++ {

		f := 0.0
		marg5d3_n = &f
		marg5d3_i = i
		marg_5d3_w = j
		macro(4)
		o := mret_5d3
		r := 0.0

		marg5d3_n = &r
		marg5d3_i = i
		marg_5d3_w = 0
		macro(4)
		q := mret_5d3

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
