package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth = 640
	screenHeight = 480
	ballSpeed = 3
	paddleSpeed = 6
)

type Object struct {
	X, Y, W, H int
}

type Paddle struct {
	Object 
}

type Ball struct {
	Object 
	dxdt int // x velocity per tick
	dydt int // y velocity per tick
}

type Game struct {
	paddle Paddle
	ball Ball
	score int
	highScore int
}

func main() {
	ebiten.SetWindowTitle("go-pong")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &Game{}

	err := ebiten.RunGame(g)

	if err != nil {
		log.Fatal(err)
	}
}

func (g*Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Update() error {
	return nil
}
