package sort_test

import (
	"fmt"
//	"math"
//	"math/rand"
	"sort"
//	"strconv"
	"testing"
	"unsafe"
	"reflect"
)
////////////////////////////////////////////////////////////////////////////////

// Sort macro
func Sort(ts0 *[]int, s reflect.SliceHeader, comparer Comparerer) {

//	fmt.Println("HELLOU", s.Len)
	compar := comparer.Comparer(ts0)
	for i := 0; i < s.Len; i++ {
		for j := 0; j < i; j++ {
			slicei := unsafe.Pointer(s.Data + uintptr(i*(*ts0)[0]))
			slicej := unsafe.Pointer(s.Data + uintptr(j*(*ts0)[0]))
			if compar(slicei, slicej) > 0 {
//				fmt.Println("hi")
//				fmt.Println((*[]float64)(unsafe.Pointer(&s)))
				for q := 0; q < (*ts0)[0]; q++ {
					sli := (*byte)(unsafe.Pointer(uintptr(slicei) + uintptr(q)))
					slj := (*byte)(unsafe.Pointer(uintptr(slicej) + uintptr(q)))
					/*
						x := (*slice)[i]
						(*slice)[i] = (*slice)[j]
						(*slice)[j] = x
					*/
					x := *sli
					*sli = *slj
					*slj = x
				}
//				fmt.Println((*[]float64)(unsafe.Pointer(&s)))
			} else {
//				fmt.Println("ho")
			}

		}
	}

}
// Comparerer macro-interface
type Comparerer interface {
	Comparer(m *[]int) func(a, b unsafe.Pointer) int
}

type Comparert struct{}

// Type specific compar functions
func Compar64(a, b *int64) int {
	return int(*a - *b)
}

// Comparer macro
func (Comparert) Comparer(ts0 *[]int) func(a, b unsafe.Pointer) int {
	switch (*ts0)[0] {

	case 8:
		return func(a, b unsafe.Pointer) int { return Compar64((*int64)(a), (*int64)(b)) }
	default:
		panic("")
	}
}
////////////////////////////////////////////////////////////////////////////////

   	type int64slice []int64
   	
   	func (p int64slice) Len() int           { return len(p) }
   	func (p int64slice) Less(i, j int) bool { return p[i] < p[j]  }
   	func (p int64slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
////////////////////////////////////////////////////////////////////////////////
func random(prng *[2]uint64) uint64 {
	s1 := prng[0]
	s0 := prng[1]
	prng[0] = s0
	s1 ^= s1 << 23 // a
	prng[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26))
	return prng[1] + s0 // b, c
}

func fillu(data []int64, seed *[2]uint64) {
	for i := 0; i < len(data); i++ {
		data[i] = int64(random(seed))
	}
}

const datasize = 8

/*
func fillf(data []float64, seed *[2]uint64) {
	for i := 0; i < len(data); i++ {
		data[i] = float64(random(seed))
	}
}
*/
func BenchmarkSortLarge_Random(b *testing.B) {
	b.StopTimer()
	var seed = [2]uint64{0x1337, 0xbeef}
	n := datasize
	if testing.Short() {
		n /= 100
	}
	data := make([]int64, n)
	fillu(data, &seed)

	fmt.Println(data)

	if sort.IsSorted(int64slice(data)) {
		b.Fatalf("terrible rand.rand")
	}
	b.StartTimer()
	sort.Sort(int64slice(data))
	b.StopTimer()
	if !sort.IsSorted(int64slice(data)) {
		b.Errorf("sort didn't sort - 1M ints")
	}
}

func BenchmarkMySortLarge_Random(b *testing.B) {
	b.StopTimer()
	var seed = [2]uint64{0x1337, 0xbeef}
	n := datasize
	if testing.Short() {
		n /= 100
	}
	data := make([]int64, n)
	fillu(data, &seed)

	fmt.Println(data)

	if sort.IsSorted(int64slice(data)) {
		b.Fatalf("terrible rand.rand")
	}
///////////
	var m []int
	m = append(m, int(unsafe.Sizeof(data[0])))

	// run the sorts

	Sort(&m, *(*reflect.SliceHeader)((unsafe.Pointer(&data))), Comparert{})

////////
	b.StartTimer()
	sort.Sort(int64slice(data))
	b.StopTimer()
	if !sort.IsSorted(int64slice(data)) {
		b.Errorf("sort didn't sort - 1M ints")
	}
}
