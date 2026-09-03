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

// Called 60 times a second
func (g *Game) Update() error {
	g.paddle.MoveOnKeyPress()
	g.ball.Move()
	g.CollideWithWall()
	g.CollideWithPaddle()
	return nil
}

// Move paddle up and down with arrow keys
func (p *Paddle) MoveOnKeyPress() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += paddleSpeed
	} 

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= paddleSpeed
	} 
}

// Ball moves on its own
func (b *Ball) Move() {
	b.X += b.dxdt
	b.Y += b.dydt
}

// Reset score and ball position
func (g *Game) Reset() {
	g.ball.X = 0
	g.ball.Y = 0
	g.score = 0
}

func (g *Game) CollideWithWall() {
	// Right wall causes game over
	if g.ball.X >= screenWidth {
		g.Reset()
	} else if g.ball.X <= 0 { // Hit left wall
		// Reverse ball's X speed and direction
		g.ball.dxdt = ballSpeed
	} else if g.ball.Y <= 0 { // Hit top wall
		// Reverse ball's Y speed and direction
		g.ball.dydt = ballSpeed
	} else if g.ball.Y <= screenHeight { // Hit bottom wall
		// Reverse ball's Y speed and direction
		g.ball.dydt = ballSpeed
	}
}

func (g *Game) CollideWithPaddle() {
	// Check that ball is in front of paddle and within its dimensions
	if g.ball.X >= g.paddle.X && g.ball.Y >= g.paddle.Y && g.ball.Y <= g.paddle.Y + g.paddle.H {
		// Reverse X speed to simulate bounce off the paddle
		g.ball.dxdt = -g.ball.dxdt

		// Add point to score
		g.score++

		// Update high score if applicable
		if g.score > g.highScore {
			g.highScore = g.score
		}
	}
}
	
