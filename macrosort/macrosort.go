package main

import "unsafe"
import "reflect"
import "math"
import "fmt"

// memsizetyper provides sizes and (optionally) types to macros
type memsizetyper struct {
	siz int
}

// Type specific compar functions
func Compar32(a, b *float32) int {
	oa := int(math.Float32bits(*a))
	ob := int(math.Float32bits(*b))
	return oa - ob
}

// Type specific compar functions
func Compar64(a, b *float64) int {
	oa := int(math.Float64bits(*a))
	ob := int(math.Float64bits(*b))
	return int(oa - ob)
}

type Comparer interface {
	Comparer(m *memsizetyper) func(a, b unsafe.Pointer) int
}

// Comparer macro
func Comparer(m *memsizetyper) func(a, b unsafe.Pointer) int {
	switch m.siz {
	case 4:
		return func(a, b unsafe.Pointer) int { return Compar32((*float32)(a), (*float32)(b)) }
	case 8:
		return func(a, b unsafe.Pointer) int { return Compar64((*float64)(a), (*float64)(b)) }
	default:
		panic("")
	}
}

// Sort macro
func Sort(s reflect.SliceHeader, comparer func(m *memsizetyper) func(a, b unsafe.Pointer) int, m *memsizetyper) {
	compar := comparer(m)
	for i := 0; i < s.Len; i++ {
		for j := 0; j < i; j++ {
			slicei := unsafe.Pointer(s.Data + uintptr(i*m.siz))
			slicej := unsafe.Pointer(s.Data + uintptr(j*m.siz))
			if compar(slicei, slicej) > 0 {
				fmt.Println((*[]float32)(unsafe.Pointer(&s)))
				for q := 0; q < m.siz; q++ {
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
				fmt.Println((*[]float32)(unsafe.Pointer(&s)))
			}

		}
	}

}

func main() {
	var sa = []float32{818, 128, 39, 153, 643}
	var sb = []float32{9, 8, 7, 6, 5, 4, 3, 1, 2}
	_ = sa
	_ = sb

	var m memsizetyper
	m.siz = int(unsafe.Sizeof(sa[0]))

	Sort(*(*reflect.SliceHeader)((unsafe.Pointer(&sa))), Comparer, &m)

	fmt.Println(sa)
}
