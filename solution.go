package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

type SideSquare int

const (
	SidesCircle = 0
	SidesTriangle = 3
	SidesSquare = 4
)

func CalcSquare(sideLen float64, sidesNum SideSquare) float64 {
	var res float64 = 0
	switch sidesNum {
	case SidesCircle:
		res = math.Pi * float64(sideLen) * float64(sideLen)
	case SidesTriangle:
		res = math.Sqrt(3) / 4 * float64(sideLen) * float64(sideLen)
	case SidesSquare:
		res = float64(sideLen) * float64(sideLen)
	}
	return res
}
