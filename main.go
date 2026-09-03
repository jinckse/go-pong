package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"golang.org/x/image/font/basicfont"
	"fmt"
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

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw paddle
	vector.DrawFilledRect(screen,
	float32(g.paddle.X), float32(g.paddle.Y),
	float32(g.paddle.W), float32(g.paddle.H),
	color.White, false)

	// Draw ball
	vector.DrawFilledRect(screen,
	float32(g.ball.X), float32(g.ball.Y),
	float32(g.ball.W), float32(g.ball.H),
	color.White, false)

	// Draw current score
	scoreStr := "Score: " + fmt.Sprint(g.score)    //X, Y
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 20, color.White)

	// Draw high score
	highScoreStr := "High Score: " + fmt.Sprint(g.highScore) //X, Y
	text.Draw(screen, highScoreStr, basicfont.Face7x13, 10, 40, color.White)
}

func (g *Game) Update() error {
	return nil
}
