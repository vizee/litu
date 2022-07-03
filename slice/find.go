package slice

import (
	"bytes"
	"litu/cmp"
	"litu/constraints"
	"litu/ptr"
	"unsafe"
)

type Position int

func (p Position) Inner() int {
	return int(p)
}

func (p Position) IsSome() bool {
	return p >= 0
}

func Find[T comparable](haystack []T, needle *T) Position {
	switch any((*T)(nil)).(type) {
	case *byte:
		// fuck type
		haystack_ := ptr.AsSlice[byte](unsafe.Pointer(ptr.FromSlice(haystack)), len(haystack))
		needle_ := ptr.Read[byte](unsafe.Pointer(&needle))
		return Position(bytes.IndexByte(haystack_, needle_))
	default:
		for i := 0; i < len(haystack); i++ {
			if haystack[i] == *needle {
				return Position(i)
			}
		}
		return Position(-1)
	}
}

func Contains[T comparable](haystack []T, needle *T) bool {
	return Find(haystack, needle).IsSome()
}

func FindBy[T any](haystack []T, needle *T, eq cmp.Equal[*T]) Position {
	for i := 0; i < len(haystack); i++ {
		if eq(&haystack[i], needle) {
			return Position(i)
		}
	}
	return Position(-1)
}

func BinarySearch[T constraints.Compare](haystack []T, needle T) Position {
	left := 0
	right := len(haystack)
	for left < right {
		mid := int(uint(left+right) / 2)
		t := haystack[mid]
		if t < needle {
			left = mid + 1
		} else if t > needle {
			right = mid
		} else {
			return Position(mid)
		}
	}
	return Position(-left - 1)
}

func BinarySearchBy[T any](haystack []T, needle *T, cmp cmp.Cmp[*T]) Position {
	left := 0
	right := len(haystack)
	for left < right {
		mid := int(uint(left+right) / 2)
		ord := cmp(&haystack[mid], needle)
		if ord < 0 {
			left = mid + 1
		} else if ord > 0 {
			right = mid
		} else {
			return Position(mid)
		}
	}
	return Position(-left - 1)
}
