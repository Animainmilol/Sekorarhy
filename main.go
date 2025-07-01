package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	InitWindowWidth  = 800
	InitWindowHeight = 450

	MaxZoom       = 15
	MinZoom       = -0.9
	RotationSpeed = 10
	ZoomSpeed     = 2
	CameraSpeed   = 8
)

type CameraController struct {
	camera         rl.Camera2D
	manualOffset   rl.Vector2
	manualRotation float32
	manualZoom     float32
	speed          float32
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

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		cameraController.updateCamera()

		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}
}
