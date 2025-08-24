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
	GridPos          rl.Vector2
	Size             float32
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
		GridPos:          rl.NewVector2(0, 0),
		Size:             SquareSize,
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
		sc.GridPos.Y -= sc.TeleportDistance
	case 'a':
		sc.GridPos.X -= sc.TeleportDistance
	case 's':
		sc.GridPos.Y += sc.TeleportDistance
	case 'd':
		sc.GridPos.X += sc.TeleportDistance
	case 'W':
		sc.GridPos.Y -= sc.TeleportDistance * LongTeleportMultiplier
	case 'A':
		sc.GridPos.X -= sc.TeleportDistance * LongTeleportMultiplier
	case 'S':
		sc.GridPos.Y += sc.TeleportDistance * LongTeleportMultiplier
	case 'D':
		sc.GridPos.X += sc.TeleportDistance * LongTeleportMultiplier
	}
}

func (sc SquareController) GetDrawPosition() rl.Vector2 {
	return rl.Vector2{
		X: sc.GridPos.X * sc.Size,
		Y: sc.GridPos.Y * sc.Size,
	}
}

func (sc SquareController) Draw(w World) {
	drawPos := sc.GetDrawPosition()
	rl.DrawRectangleV(drawPos, rl.NewVector2(sc.Size, sc.Size), sc.Color)
}

func (sc SquareController) GetCenter() rl.Vector2 {
	drawPos := sc.GetDrawPosition()
	return rl.Vector2{
		X: drawPos.X + sc.Size/2,
		Y: drawPos.Y + sc.Size/2,
	}
}
