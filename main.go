package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	InitWindowWidth  = 800
	InitWindowHeight = 450

	SquareSize             = 50
	TeleportDistance       = 100
	LongTeleportMultiplier = 2

	MaxZoom       = 15
	MinZoom       = -0.9
	RotationSpeed = 10
	ZoomSpeed     = 2
	CameraSpeed   = 8
)

type SquareController struct {
	rectangle        rl.Rectangle
	color            rl.Color
	teleportDistance float32
}

type CameraController struct {
	camera         rl.Camera2D
	manualOffset   rl.Vector2
	manualRotation float32
	manualZoom     float32
	speed          float32
}

type World struct {
	tiles map[[2]int32]rl.Color
}

func (sc *SquareController) handleSquareMovementInput() {
	if rl.IsKeyPressed(rl.KeyRight) {
		sc.rectangle.X += sc.teleportDistance
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		sc.rectangle.X -= sc.teleportDistance
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		sc.rectangle.Y += sc.teleportDistance
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		sc.rectangle.Y -= sc.teleportDistance
	}

	if rl.IsKeyPressed(rl.KeyD) {
		sc.rectangle.X += sc.teleportDistance * LongTeleportMultiplier
	}
	if rl.IsKeyPressed(rl.KeyA) {
		sc.rectangle.X -= sc.teleportDistance * LongTeleportMultiplier
	}
	if rl.IsKeyPressed(rl.KeyS) {
		sc.rectangle.Y += sc.teleportDistance * LongTeleportMultiplier
	}
	if rl.IsKeyPressed(rl.KeyW) {
		sc.rectangle.Y -= sc.teleportDistance * LongTeleportMultiplier
	}
}

func (cc *CameraController) handleCameraControlInput() {
	if rl.IsKeyDown(rl.KeyQ) {
		cc.manualRotation += RotationSpeed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyE) {
		cc.manualRotation -= RotationSpeed * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyZ) {
		cc.manualZoom = min(cc.manualZoom+ZoomSpeed*rl.GetFrameTime(), MaxZoom)
	}
	if rl.IsKeyDown(rl.KeyX) {
		cc.manualZoom = max(cc.manualZoom-ZoomSpeed*rl.GetFrameTime(), MinZoom)
	}
}

func drawWorld(w World, cc CameraController) {
	screenWidth := float32(rl.GetRenderWidth())
	screenHeight := float32(rl.GetRenderHeight())

	corners := [4]rl.Vector2{
		rl.GetScreenToWorld2D(rl.NewVector2(0, 0), cc.camera),                      // top-left
		rl.GetScreenToWorld2D(rl.NewVector2(screenWidth, 0), cc.camera),            // top-right
		rl.GetScreenToWorld2D(rl.NewVector2(screenWidth, screenHeight), cc.camera), // bottom-right
		rl.GetScreenToWorld2D(rl.NewVector2(0, screenHeight), cc.camera),           // bottom-left
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

	// Convert to tile coordinates (2 is padding)
	startX := int32(min(minX)/SquareSize) - 2
	startY := int32(min(minY)/SquareSize) - 2
	endX := int32(max(maxX)/SquareSize) + 2
	endY := int32(max(maxY)/SquareSize) + 2

	visibleTiles := make(map[[2]int32]rl.Color)

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			color, exists := w.tiles[[2]int32{x, y}]
			if exists {
				visibleTiles[[2]int32{x, y}] = color
			}
		}
	}

	for coordinates, color := range visibleTiles {
		x, y := coordinates[0], coordinates[1]
		rl.DrawRectangle(x*SquareSize, y*SquareSize, SquareSize, SquareSize, color)
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

	cc.camera.Offset = rl.Vector2Add(centeringOffset, cc.manualOffset)
	cc.camera.Zoom = scalingZoom + cc.manualZoom*scale
	cc.camera.Rotation = cc.manualRotation
}

func cameraFollow(cc *CameraController, x float32, y float32) {
	cameraSpeed := cc.speed * rl.GetFrameTime()
	cc.camera.Target.X += (x - cc.camera.Target.X) * cameraSpeed
	cc.camera.Target.Y += (y - cc.camera.Target.Y) * cameraSpeed
}

func newSquareController() *SquareController {
	return &SquareController{
		rectangle:        rl.NewRectangle(0, 0, SquareSize, SquareSize),
		color:            rl.White,
		teleportDistance: TeleportDistance,
	}
}

func newCameraController() *CameraController {
	return &CameraController{
		camera: rl.Camera2D{
			Offset:   rl.NewVector2(0, 0),
			Target:   rl.NewVector2(0, 0),
			Rotation: 0,
			Zoom:     1,
		},
		speed: CameraSpeed,
	}
}

func main() {
	rl.InitWindow(InitWindowWidth, InitWindowHeight, "Sekorathy")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	cameraController := newCameraController()
	squareController := newSquareController()

	world := World{
		tiles: make(map[[2]int32]rl.Color),
	}

	// test
	gridSize := int32(1000)
	for y := int32(0); y < gridSize; y++ {
		for x := int32(0); x < gridSize; x++ {
			world.tiles[[2]int32{x, y}] = rl.Red
		}
	}

	for !rl.WindowShouldClose() {
		squareController.handleSquareMovementInput()
		cameraController.handleCameraControlInput()

		rl.BeginDrawing()

		cameraController.updateCamera()
		cameraFollow(
			cameraController,
			squareController.rectangle.X+squareController.rectangle.Width/2,
			squareController.rectangle.Y+squareController.rectangle.Height/2,
		)

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(cameraController.camera)

		drawWorld(world, *cameraController)
		rl.DrawRectangleRec(squareController.rectangle, squareController.color)

		rl.EndMode2D()

		rl.EndDrawing()
	}
}
