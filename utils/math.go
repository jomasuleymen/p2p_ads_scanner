package utils

import (
	"math"
	"math/rand"
	"strconv"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func ParseFloat(str string) float64 {
	value, _ := strconv.ParseFloat(str, 64)
	return value
}

func RandomInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
