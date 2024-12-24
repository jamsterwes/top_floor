package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Graph struct {
	samples      []float32
	bufferSize   int
	nextSample   int
	cachedPixels *[]byte
}

func createGraph(bufferSize int) Graph {
	// Initialize the samples to [0, 0, ...]
	graph := Graph{}
	graph.samples = make([]float32, bufferSize)
	graph.bufferSize = bufferSize
	graph.nextSample = 0
	graph.cachedPixels = nil

	// Return a graph
	return graph
}

// Add a new sample to the graph
func (graph *Graph) AddSample(x float32) {
	if graph.nextSample < graph.bufferSize {
		// If nextSample < bufferSize, just add in place
		graph.samples[graph.nextSample] = x
		graph.nextSample += 1
	} else {
		// Otherwise, push all samples back and insert
		graph.samples = append(graph.samples[1:], x)
	}
}

// Draw the graph (assume 640 samples, $0 to $479 price)
func (graph *Graph) Draw(screen *ebiten.Image) {
	// TODO: change size
	img := make([]byte, 640*480*4)
	graph.cachedPixels = &img

	// Now draw the graph
	for x := range graph.nextSample {
		// If x == 0, last == this
		lastX := x
		if x > 0 {
			lastX -= 1
		}

		lastY := 479 - int(min(479, graph.samples[lastX]))
		currY := 479 - int(min(479, graph.samples[x]))

		colorChannel := 1
		if currY > lastY {
			colorChannel = 0
		}

		for y := min(lastY, currY); y <= max(lastY, currY); y++ {
			(*graph.cachedPixels)[y*640*4+x*4+colorChannel] = 0xFF
			(*graph.cachedPixels)[y*640*4+x*4+3] = 0xFF
		}
	}
	screen.WritePixels(*graph.cachedPixels)
}
