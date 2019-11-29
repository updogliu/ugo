package ugo

import "sort"

func AppendStr(a *[]string, val string) {
	*a = append(*a, val)
}

func AppendInt(a *[]int, val int) {
	*a = append(*a, val)
}

func AppendI64(a *[]int64, val int64) {
	*a = append(*a, val)
}

func AppendUint(a *[]uint, val uint) {
	*a = append(*a, val)
}

func AppendU8(a *[]uint8, val uint8) {
	*a = append(*a, val)
}

func AppendU64(a *[]uint64, val uint64) {
	*a = append(*a, val)
}

func AppendByte(a *[]byte, val byte) {
	*a = append(*a, val)
}

func AppendF64(a *[]float64, val float64) {
	*a = append(*a, val)
}

func SortI64(a []int64) {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
}

func SortU64(a []uint64) {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
}
