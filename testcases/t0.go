// +build gomacro_1

package main

// begin macro
func Compar(a, b *) int {
	return int(*a) - int(*b)
}
// end macro

func check(i,v int) (r bool) {
	if v != 0 && i&3 == 0 {
		print("00:")
		r = true
	}
	if v != 0 && i&3 == 3 {
		print("11:")
		r = true
	}
	if v >= 0 && i&3 == 1 {
		print("01:")
		r = true
	}
	if v <= 0 && i&3 == 2 {
		print("10:")
		r = true
	}
	if r {
		print(i)
	}
	return r
}

func main() {
	var r [1024]int

	var v0, v1 int = 3,4
	var w0, w1 int = 3,4
	var q0, q1 int = 3,4
	var p0, p1 int = 3,4
	var o0, o1 int = 3,4

	var i = 0

	// begin tests

	r[i|0] = Compar(&v0, &v0)
	r[i|1] = Compar(&v0, &v1)
	r[i|2] = Compar(&v1, &v0)
	r[i|3] = Compar(&v1, &v1)

	i += 4

	r[i|0] = Compar(&w0, &w0)
	r[i|1] = Compar(&w0, &w1)
	r[i|2] = Compar(&w1, &w0)
	r[i|3] = Compar(&w1, &w1)

	i += 4

	r[i|0] = Compar(&q0, &q0)
	r[i|1] = Compar(&q0, &q1)
	r[i|2] = Compar(&q1, &q0)
	r[i|3] = Compar(&q1, &q1)

	i += 4

	r[i|0] = Compar(&p0, &p0)
	r[i|1] = Compar(&p0, &p1)
	r[i|2] = Compar(&p1, &p0)
	r[i|3] = Compar(&p1, &p1)

	i += 4

	r[i|0] = Compar(&o0, &o0)
	r[i|1] = Compar(&o0, &o1)
	r[i|2] = Compar(&o1, &o0)
	r[i|3] = Compar(&o1, &o1)

	i += 4

	// end tests

	for j := 0; j < i; j++ {
		if check(j,r[j]) {
			panic("!!!")
		}
	}
}
