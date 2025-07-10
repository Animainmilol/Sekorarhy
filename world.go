package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WorldRenderPadding = 2
)

type World struct {
	tiles map[[2]int32]Tile
}

type Tile struct {
	Type string
}

func (w World) Draw(cc CameraController) {
	screenWidth := float32(rl.GetRenderWidth())
	screenHeight := float32(rl.GetRenderHeight())

	corners := [4]rl.Vector2{
		rl.GetScreenToWorld2D(rl.NewVector2(0, 0), cc.Camera),                      // top-left
		rl.GetScreenToWorld2D(rl.NewVector2(screenWidth, 0), cc.Camera),            // top-right
		rl.GetScreenToWorld2D(rl.NewVector2(screenWidth, screenHeight), cc.Camera), // bottom-right
		rl.GetScreenToWorld2D(rl.NewVector2(0, screenHeight), cc.Camera),           // bottom-left
	}

	minX := corners[0].X
	maxX := corners[0].X
	minY := corners[0].Y
	maxY := corners[0].Y

	for _, corner := range corners[1:] {
		minX = min(minX, corner.X)
		maxX = max(maxX, corner.X)
		minY = min(minY, corner.Y)
		maxY = max(maxY, corner.Y)
	}

	// Convert to tile coordinates
	startX := int32(minX/SquareSize) - WorldRenderPadding
	startY := int32(minY/SquareSize) - WorldRenderPadding
	endX := int32(maxX/SquareSize) + WorldRenderPadding
	endY := int32(maxY/SquareSize) + WorldRenderPadding

	visibleTiles := make(map[[2]int32]Tile)

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			tile, exists := w.tiles[[2]int32{x, y}]
			if exists {
				visibleTiles[[2]int32{x, y}] = tile
			}
		}
	}

	for coordinates, tile := range visibleTiles {
		x, y := coordinates[0], coordinates[1]
		switch tile.Type {
		case "dot":
			rl.DrawCircle(x*SquareSize+SquareSize/2, y*SquareSize+SquareSize/2, SquareSize/4, rl.White)
		case "box":
			rl.DrawRectangle(x*SquareSize, y*SquareSize+20, SquareSize, SquareSize-40, rl.White)
		default:
			rl.DrawRectangle(x*SquareSize, y*SquareSize, SquareSize, SquareSize, rl.Red)
		}
	}
}

func placeTilesUsingCursor(w World, cc CameraController) {
	// Calculate world coordinates of cursor
	cursorelPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), cc.Camera)

	// Floor is for negatives
	y := int32(math.Floor(float64(cursorelPos.Y / SquareSize)))
	x := int32(math.Floor(float64(cursorelPos.X / SquareSize)))

	if rl.IsKeyDown(rl.KeyC) {
		w.tiles[[2]int32{x, y}] = Tile{"dot"}
	}
	if rl.IsKeyDown(rl.KeyV) {
		w.tiles[[2]int32{x, y}] = Tile{"box"}
	}
}

func (w World) buildMap(movementLine string) {
	var realPos rl.Vector2
	var relPos rl.Vector2
	w.tiles[[2]int32{0, 0}] = Tile{"dot"}
	for _, r := range movementLine {
		switch r {
		case 'w':
			relPos = rl.Vector2{X: 0, Y: -1}
		case 'a':
			relPos = rl.Vector2{X: -1, Y: 0}
		case 's':
			relPos = rl.Vector2{X: 0, Y: 1}
		case 'd':
			relPos = rl.Vector2{X: 1, Y: 0}
		case 'W':
			relPos = rl.Vector2{X: 0, Y: -2}
		case 'A':
			relPos = rl.Vector2{X: -2, Y: 0}
		case 'S':
			relPos = rl.Vector2{X: 0, Y: 2}
		case 'D':
			relPos = rl.Vector2{X: 2, Y: 0}
		}
		realPos.X += relPos.X * 2
		realPos.Y += relPos.Y * 2
		x := int32(realPos.X)
		y := int32(realPos.Y)
		w.tiles[[2]int32{x, y}] = Tile{"dot"}
	}
}
