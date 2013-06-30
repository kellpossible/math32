package math32

import (
	"math"
)

const Pi = math.Pi

func Max(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

func MaxSlice(slice []float32) float32 {
	var max float32 = -1.0 * (math.MaxFloat32 - 100.0)
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}

func Abs(a float32) float32 {
	return float32(math.Abs(float64(a)))
}

func Sqrt(a float32) float32 {
	return float32(math.Sqrt(float64(a)))
}

func Sin(a float32) float32 {
	return float32(math.Sin(float64(a)))
}

func Cos(a float32) float32 {
	return float32(math.Cos(float64(a)))
}

func Tan(a float32) float32 {
	return float32(math.Tan(float64(a)))
}

func Asin(a float32) float32 {
	return float32(math.Asin(float64(a)))
}

func Acos(a float32) float32 {
	return float32(math.Acos(float64(a)))
}

func Atan(a float32) float32 {
	return float32(math.Atan(float64(a)))
}

func Pow(a, b float32) float32 {
	return float32(math.Pow(float64(a), float64(b)))
}

func Pow10(a int) float32 {
	return float32(math.Pow10(a))
}

func Bool2Float(a bool) float32 {
	if a {
		return 1.0
	}
	return 0.0
}

func Xor(x, y bool) bool {
	switch {
	case x && !y:
		return true
	case y && !x:
		return true
	}
	return false
}
