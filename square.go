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

func (sc *SquareController) HandleInput(movementLine string) {
	movement := getCurrentMovement()
	if movement == 0 {
		return
	}

	if sc.isCorrectMovement(movementLine, movement) {
		sc.executeMovement(movement)
	} else {
		if sc.hasMovementsLeft(movementLine) {
			current := movementLine[sc.Step]
			sc.executeMovement(rune(current))
		}
	}
}

func getCurrentMovement() rune {
    for key, movement := range keyBindings {
        if rl.IsKeyPressed(key) {
            return movement
        }
    }
    return 0
}

func (sc *SquareController) hasMovementsLeft(movementLine string) bool {
	return sc.Step < int32(len(movementLine))
}

func (sc *SquareController) isCorrectMovement(movementLine string, movement rune) bool {
	if !sc.hasMovementsLeft(movementLine) {
		return false
	}

	current := movementLine[sc.Step]
	if current == '/' {
		sc.Step++
	}
	return current == byte(movement)
}

func (sc *SquareController) executeMovement(movement rune) {
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
