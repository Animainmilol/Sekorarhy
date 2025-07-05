package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SquareSize             = 50
	TeleportDistance       = 100
	LongTeleportMultiplier = 2
)

type SquareController struct {
	Rectangle        rl.Rectangle
	Color            rl.Color
	TeleportDistance float32
}

func NewSquareController() *SquareController {
	return &SquareController{
		Rectangle:        rl.NewRectangle(0, 0, SquareSize, SquareSize),
		Color:            rl.White,
		TeleportDistance: TeleportDistance,
	}
}

func (sc *SquareController) HandleInput(s rl.Sound) {
	if rl.IsKeyPressed(rl.KeyRight) {
		sc.Rectangle.X += sc.TeleportDistance
		rl.PlaySound(s)
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		sc.Rectangle.X -= sc.TeleportDistance
		rl.PlaySound(s)
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		sc.Rectangle.Y += sc.TeleportDistance
		rl.PlaySound(s)
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		sc.Rectangle.Y -= sc.TeleportDistance
		rl.PlaySound(s)
	}

	if rl.IsKeyPressed(rl.KeyD) {
		sc.Rectangle.X += sc.TeleportDistance * LongTeleportMultiplier
		rl.PlaySound(s)
	}
	if rl.IsKeyPressed(rl.KeyA) {
		sc.Rectangle.X -= sc.TeleportDistance * LongTeleportMultiplier
		rl.PlaySound(s)
	}
	if rl.IsKeyPressed(rl.KeyS) {
		sc.Rectangle.Y += sc.TeleportDistance * LongTeleportMultiplier
		rl.PlaySound(s)
	}
	if rl.IsKeyPressed(rl.KeyW) {
		sc.Rectangle.Y -= sc.TeleportDistance * LongTeleportMultiplier
		rl.PlaySound(s)
	}
}

func recordMovement() string {
	var movement string
	if rl.IsKeyPressed(rl.KeyRight) {
		movement = "d"
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		movement = "a"
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		movement = "s"
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		movement = "w"
	}

	if rl.IsKeyPressed(rl.KeyD) {
		movement = "D"
	}
	if rl.IsKeyPressed(rl.KeyA) {
		movement = "A"
	}
	if rl.IsKeyPressed(rl.KeyS) {
		movement = "S"
	}
	if rl.IsKeyPressed(rl.KeyW) {
		movement = "W"
	}

	return movement
}

func (sc SquareController) GetCenter() rl.Vector2 {
	return rl.Vector2{
		X: sc.Rectangle.X + sc.Rectangle.Width/2,
		Y: sc.Rectangle.Y + sc.Rectangle.Height/2,
	}
}
