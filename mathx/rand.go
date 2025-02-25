package mathx

import "math/rand"

func RandInRange[T int | int8 | int16 | int32 | int64](min, max T) T {
	return T(rand.Int63n(int64(max-min+1))) + min
}
