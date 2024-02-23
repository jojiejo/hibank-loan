package helpers

import "math"

func Round(value float64, decimalPoints int) float64 {
	factor := math.Pow10(decimalPoints)
	rounded := math.Round(value * factor)
	return rounded / factor
}
