package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	MaxZoom       = 15
	MinZoom       = -0.9
	RotationSpeed = 30
	ZoomSpeed     = 2
	CameraSpeed   = 8
)

type CameraController struct {
	Camera         rl.Camera2D
	ManualOffset   rl.Vector2
	ManualRotation float32
	ManualZoom     float32
	Speed          float32
}

func NewCameraController() *CameraController {
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

func (cc *CameraController) HandleInput() {
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

func (cc *CameraController) Update() {
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

func CameraFollow(cc *CameraController, x float32, y float32) {
	cameraSpeed := cc.Speed * rl.GetFrameTime()
	cc.Camera.Target.X += (x - cc.Camera.Target.X) * cameraSpeed
	cc.Camera.Target.Y += (y - cc.Camera.Target.Y) * cameraSpeed
}
