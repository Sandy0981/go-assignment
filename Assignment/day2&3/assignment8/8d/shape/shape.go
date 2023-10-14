package shape

import (
	"8d/model"
	"math"
)

func Area(a model.Circle) float32 {
	return math.Pi * a.Radius * a.Radius
}
