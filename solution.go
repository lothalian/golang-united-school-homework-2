package square

import (
	"math"
)

type sides int

const (
	SidesTriangle sides = 3
	SidesSquare   sides = 4
	SidesCircle   sides = 0
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {

	var area float64

	if sideLen <= 0 {
		return 0
	} else {
		switch {
		case sidesNum == 3:
			area = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
		case sidesNum == 4:
			area = math.Pow(sideLen, 2)
		case sidesNum == 0:
			area = math.Pow(sideLen, 2) * math.Pi
		default:
			return 0
		}
		return area
	}
}
