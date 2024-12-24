package main

import (
	"math/rand"
)

func randomRangeFloat32(low float32, high float32) float32 {
	return rand.Float32()*(high-low) + low
}
