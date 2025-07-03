package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	InitWindowWidth  = 800
	InitWindowHeight = 450

	SquareSize             = 50
	TeleportDistance       = 100
	LongTeleportMultiplier = 2

	MaxZoom            = 15
	MinZoom            = -0.9
	RotationSpeed      = 30
	ZoomSpeed          = 2
	CameraSpeed        = 8
	WorldRenderPadding = 2
)

type SquareController struct {
	Rectangle        rl.Rectangle
	Color            rl.Color
	TeleportDistance float32
}

type CameraController struct {
	Camera         rl.Camera2D
	ManualOffset   rl.Vector2
	ManualRotation float32
	ManualZoom     float32
	Speed          float32
}

type World struct {
	tiles map[[2]int32]Tile
}

type Tile struct {
	Type string
}

func (sc *SquareController) handleInput(s rl.Sound) {
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

func (cc *CameraController) handleInput() {
	if rl.IsKeyDown(rl.KeyQ) {
		cc.ManualRotation += RotationSpeed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyE) {
		cc.ManualRotation -= RotationSpeed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyZ) {
		cc.ManualZoom = min(cc.ManualZoom+ZoomSpeed*rl.GetFrameTime(), MaxZoom)
	}
	if rl.IsKeyDown(rl.KeyX) {
		cc.ManualZoom = max(cc.ManualZoom-ZoomSpeed*rl.GetFrameTime(), MinZoom)
	}
}

func handleInput(sc *SquareController, cc *CameraController, s rl.Sound, w World) {
	sc.handleInput(s)
	cc.handleInput()
	placeTilesUsingCursor(w, *cc)
}

func drawFrame(w World, sc SquareController, cc *CameraController) {
	rl.BeginDrawing()

	cc.updateCamera()
	cameraFollow(
		cc,
		sc.getCenter().X,
		sc.getCenter().Y,
	)

	rl.ClearBackground(rl.Black)

	rl.BeginMode2D(cc.Camera)

	drawWorld(w, *cc)

	rl.DrawRectangleRec(sc.Rectangle, sc.Color)

	rl.EndMode2D()

	rl.EndDrawing()
}

func drawWorld(w World, cc CameraController) {
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
		if corner.X < minX {
			minX = corner.X
		}
		if corner.X > maxX {
			maxX = corner.X
		}
		if corner.Y < minY {
			minY = corner.Y
		}
		if corner.Y > maxY {
			maxY = corner.Y
		}
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
	cursorPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), cc.Camera)

	// Floor is for negatives
	y := int32(math.Floor(float64(cursorPos.Y / SquareSize)))
	x := int32(math.Floor(float64(cursorPos.X / SquareSize)))

	if rl.IsKeyDown(rl.KeyC) {
		w.tiles[[2]int32{x, y}] = Tile{"dot"}
	}
	if rl.IsKeyDown(rl.KeyV) {
		w.tiles[[2]int32{x, y}] = Tile{"box"}
	}
}

func (cc *CameraController) updateCamera() {
	// Center the camera
	screenWidth := float32(rl.GetRenderWidth())
	screenHeight := float32(rl.GetRenderHeight())
	centeringOffset := rl.NewVector2(screenWidth/2, screenHeight/2)

	// Zoom according to the window size
	widthScale := float32(rl.GetRenderWidth()) / float32(InitWindowWidth)
	heightScale := float32(rl.GetRenderHeight()) / float32(InitWindowHeight)
	scale := max(widthScale, heightScale)
	scalingZoom := scale

	cc.Camera.Offset = rl.Vector2Add(centeringOffset, cc.ManualOffset)
	cc.Camera.Zoom = scalingZoom + cc.ManualZoom*scale
	cc.Camera.Rotation = cc.ManualRotation
}

func cameraFollow(cc *CameraController, x float32, y float32) {
	cameraSpeed := cc.Speed * rl.GetFrameTime()
	cc.Camera.Target.X += (x - cc.Camera.Target.X) * cameraSpeed
	cc.Camera.Target.Y += (y - cc.Camera.Target.Y) * cameraSpeed
}

func newSquareController() *SquareController {
	return &SquareController{
		Rectangle:        rl.NewRectangle(0, 0, SquareSize, SquareSize),
		Color:            rl.White,
		TeleportDistance: TeleportDistance,
	}
}

func newCameraController() *CameraController {
	return &CameraController{
		Camera: rl.Camera2D{
			Offset:   rl.NewVector2(0, 0),
			Target:   rl.NewVector2(0, 0),
			Rotation: 0,
			Zoom:     1,
		},
		Speed: CameraSpeed,
	}
}

func (sc SquareController) getCenter() rl.Vector2 {
	return rl.Vector2{
		X: sc.Rectangle.X + sc.Rectangle.Width/2,
		Y: sc.Rectangle.Y + sc.Rectangle.Height/2,
	}
}

func main() {
	rl.InitWindow(InitWindowWidth, InitWindowHeight, "Sekorathy")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	cameraController := newCameraController()
	squareController := newSquareController()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	sound := rl.LoadSound("sound.wav")
	defer rl.UnloadSound(sound)

	world := World{
		tiles: make(map[[2]int32]Tile),
	}

	for !rl.WindowShouldClose() {
		handleInput(squareController, cameraController, sound, world)
		drawFrame(world, *squareController, cameraController)
	}
}
