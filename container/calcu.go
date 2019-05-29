package cal

import "math"





func CalTriangle(a,b float64) float64{
	result:=a*a+b*b
	return math.Sqrt(result)
}