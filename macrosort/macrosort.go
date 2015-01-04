/*
generic n^2 sort
Copyright (C) 2015  github.com/anlhord

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
*/

package main

import "unsafe"
import "reflect"
import "math"
import "fmt"

// Type specific compar functions
func Compar32(a, b *float32) int {
	oa := int(math.Float32bits(*a))
	ob := int(math.Float32bits(*b))
	return oa - ob
}

// Type specific compar functions
func Compar64(a, b *float64) int {
	oa := int64(math.Float64bits(*a))
	ob := int64(math.Float64bits(*b))
	ab := oa - ob
	m := int(int64(ab >> 32))
	if m == 0 {
		return int(ab)
	}
	return m
}

// Comparerer macro-interface
type Comparerer interface {
	Comparer(m *[]int) func(a, b unsafe.Pointer) int
}

type Comparert struct{}

// Comparer macro
func (Comparert) Comparer(ts0 *[]int) func(a, b unsafe.Pointer) int {
	switch (*ts0)[0] {
	case 4:
		return func(a, b unsafe.Pointer) int { return Compar32((*float32)(a), (*float32)(b)) }
	case 8:
		return func(a, b unsafe.Pointer) int { return Compar64((*float64)(a), (*float64)(b)) }
	default:
		panic("")
	}
}

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

func main() {
	var sa = []float32{818, 128, 39, 153, 643}
	var sb = []float64{456, 7894, 354, 134, 978, 354, 2254, 85, 4567}
	_ = sa
	_ = sb

	// construct the type stacks

	var m, n []int
	m = append(m, int(unsafe.Sizeof(sa[0])))
	n = append(n, int(unsafe.Sizeof(sb[0])))

	// run the sorts

	Sort(&m, *(*reflect.SliceHeader)((unsafe.Pointer(&sa))), Comparert{})
	Sort(&n, *(*reflect.SliceHeader)((unsafe.Pointer(&sb))), Comparert{})

	// print the results

	fmt.Println(sa)
	fmt.Println(sb)
}
