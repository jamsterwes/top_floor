package main

import (
	"math"
	"math/rand"
)

type Stock struct {
	currentPrice  float32
	initialDrift  float32
	initialStdDev float32
}

// Simulation tick for stock
func (stock *Stock) Simulate(dt float32) {
	// Get delta S = S * (drift * dt + stddev * N(0,1) * sqrt(dt))
	deltaS := stock.currentPrice * (stock.initialDrift*dt + stock.initialStdDev*float32(math.Sqrt(float64(dt))*rand.NormFloat64()))

	// Apply delta S
	stock.currentPrice += deltaS

	// Clamp price to $0
	stock.currentPrice = max(0, stock.currentPrice)
}
