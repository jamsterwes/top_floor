package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	stock Stock
	graph Graph
}

func Init() *Game {
	// Create an empty game
	game := Game{}

	// Generate a random stock
	game.stock = Stock{}
	game.stock.currentPrice = 50
	game.stock.initialDrift = 0.1
	game.stock.initialStdDev = 0.2

	// Create a graph
	game.graph = createGraph(640)

	// Return pointer to game
	return &game
}

func (g *Game) Update() error {
	// Simulate stock
	g.stock.Simulate(1.0 / float32(ebiten.TPS()))

	// Add sample
	g.graph.AddSample(g.stock.currentPrice)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the graph
	g.graph.Draw(screen)

	// Draw the stock price (Debug)
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Stock Price: $%0.02f", g.stock.currentPrice), 0, 0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Stocks")
	game := Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
