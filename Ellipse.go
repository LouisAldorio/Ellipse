package Ellipse

import (
	"math"
)

// Get Eccentricity of Ellipse
func GetEccentricity(A,B float64) float64 {
 	return (math.Sqrt(math.Pow(A, 2) - math.Pow(B, 2))) / A
}

