package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SquareSize             = 50
	TeleportDistance       = 100
	LongTeleportMultiplier = 2
)

type SquareController struct {
	Rectangle        rl.Rectangle
	Color            rl.Color
	TeleportDistance float32
	Step             int32
}

func NewSquareController() *SquareController {
	return &SquareController{
		Rectangle:        rl.NewRectangle(0, 0, SquareSize, SquareSize),
		Color:            rl.White,
		TeleportDistance: TeleportDistance,
	}
}

func getCurrentMovement() rune {
	var movement rune
	if rl.IsKeyPressed(rl.KeyRight) {
		movement = 'd'
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		movement = 'a'
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		movement = 's'
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		movement = 'w'
	}

	if rl.IsKeyPressed(rl.KeyD) {
		movement = 'D'
	}
	if rl.IsKeyPressed(rl.KeyA) {
		movement = 'A'
	}
	if rl.IsKeyPressed(rl.KeyS) {
		movement = 'S'
	}
	if rl.IsKeyPressed(rl.KeyW) {
		movement = 'W'
	}

	return movement
}

func (sc *SquareController) isCorrectMovement(movementLine string, movement rune) bool {
	for sc.Step < int32(len(movementLine)) {
		current := movementLine[sc.Step]
		if current == '/' {
			sc.Step++
			continue
		}
		return current == byte(movement)
	}
	return false
}

func (sc *SquareController) move(movement rune) {
	switch movement {
	case 'w':
		sc.Rectangle.Y -= sc.TeleportDistance
	case 'a':
		sc.Rectangle.X -= sc.TeleportDistance
	case 's':
		sc.Rectangle.Y += sc.TeleportDistance
	case 'd':
		sc.Rectangle.X += sc.TeleportDistance
	case 'W':
		sc.Rectangle.Y -= sc.TeleportDistance * LongTeleportMultiplier
	case 'A':
		sc.Rectangle.X -= sc.TeleportDistance * LongTeleportMultiplier
	case 'S':
		sc.Rectangle.Y += sc.TeleportDistance * LongTeleportMultiplier
	case 'D':
		sc.Rectangle.X += sc.TeleportDistance * LongTeleportMultiplier
	}
	sc.Step++
}

func (sc SquareController) GetCenter() rl.Vector2 {
	return rl.Vector2{
		X: sc.Rectangle.X + sc.Rectangle.Width/2,
		Y: sc.Rectangle.Y + sc.Rectangle.Height/2,
	}
}
