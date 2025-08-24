package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SquareSize             = 50
	TeleportDistance       = 2
	LongTeleportMultiplier = 2
)

type SquareController struct {
	Rectangle        rl.Rectangle
	Color            rl.Color
	TeleportDistance float32
	Step             int32
}

var keyBindings = map[int32]rune{
	rl.KeyRight: 'd',
	rl.KeyLeft:  'a',
	rl.KeyDown:  's',
	rl.KeyUp:    'w',
	rl.KeyD:     'D',
	rl.KeyA:     'A',
	rl.KeyS:     'S',
	rl.KeyW:     'W',
}

func NewSquareController() *SquareController {
	return &SquareController{
		Rectangle:        rl.NewRectangle(0, 0, SquareSize, SquareSize),
		Color:            rl.White,
		TeleportDistance: TeleportDistance,
	}
}

func (sc *SquareController) HandleInput() {
	movement := getCurrentMovement()
	sc.executeMovement(movement)
	sc.Step += 1
}

func getCurrentMovement() rune {
	for key, movement := range keyBindings {
		if rl.IsKeyPressed(key) {
			return movement
		}
	}
	return 0
}

func (sc *SquareController) executeMovement(movement rune) {
	switch movement {
	case 'w':
		sc.Rectangle.Y -= sc.TeleportDistance * SquareSize
	case 'a':
		sc.Rectangle.X -= sc.TeleportDistance * SquareSize
	case 's':
		sc.Rectangle.Y += sc.TeleportDistance * SquareSize
	case 'd':
		sc.Rectangle.X += sc.TeleportDistance * SquareSize
	case 'W':
		sc.Rectangle.Y -= sc.TeleportDistance * SquareSize * LongTeleportMultiplier
	case 'A':
		sc.Rectangle.X -= sc.TeleportDistance * SquareSize * LongTeleportMultiplier
	case 'S':
		sc.Rectangle.Y += sc.TeleportDistance * SquareSize * LongTeleportMultiplier
	case 'D':
		sc.Rectangle.X += sc.TeleportDistance * SquareSize * LongTeleportMultiplier
	}
}

func (sc SquareController) Draw() {
	rl.DrawRectangleRec(sc.Rectangle, sc.Color)
}

func (sc SquareController) GetCenter() rl.Vector2 {
	return rl.Vector2{
		X: sc.Rectangle.X + sc.Rectangle.Width/2,
		Y: sc.Rectangle.Y + sc.Rectangle.Height/2,
	}
}
