package constraints

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Int | Uint
}

type Number interface {
	Int | Uint | Float
}

type Compare interface {
	Number | ~string
}

type Accumulate interface {
	Number | ~string
}
