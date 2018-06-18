package agnumber

import (
	"math"
	"strconv"

	"github.com/pkg/errors"
)

// F64toa converts float64 value to 10-based string.
// Function takes optional argument - precision - which is described in strconv.FormatFloat
func F64toa(x float64, precision ...int) string {
	p := -1
	if len(precision) > 0 {
		p = precision[0]
	}
	return strconv.FormatFloat(x, 'f', p, 64)
}

// Atoi64 converts 10-based string into int64 value.
func Atoi64(s string) (int64, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	return n, errors.Wrap(err, "can't parse int")
}

// Atof64 converts 10-based string into float64 value.
func Atof64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	return f, errors.Wrap(err, "can't parse float")
}

// I64toa converts int64 value to 10-based string
func I64toa(x int64) string {
	return strconv.FormatInt(x, 10)
}

// I64tox converts int64 value to 16-based string
func I64tox(x int64) string {
	return strconv.FormatInt(x, 16)
}

func reduceIntSlice(a []int, start int, reduce func(float64, float64) float64) int {
	if len(a) == 0 {
		return start
	}
	result := a[0]
	for _, v := range a {
		result = int(reduce(float64(result), float64(v)))
	}
	return result
}

// Min finds the minimum value in int slice
func Min(values ...int) int {
	return reduceIntSlice(values, int(math.MinInt64), math.Min)
}

// Max finds the maximum value in int slice
func Max(values ...int) int {
	return reduceIntSlice(values, int(math.MaxInt64), math.Max)
}
