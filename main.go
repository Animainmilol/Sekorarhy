package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	InitWindowWidth  = 800
	InitWindowHeight = 450

	MovementLine = "/AAWDWWAAwwdDwWwaAaWwwDddSdDwWwwwDddSsssDdwwWaaAAAaaAAa"
)

func handleInput(sc *SquareController, cc *CameraController, s rl.Sound, w World) {
	sc.HandleInput(s)
	cc.HandleInput()
	placeTilesUsingCursor(w, *cc)
}

func drawFrame(w World, sc SquareController, cc *CameraController) {
	rl.BeginDrawing()

	cc.UpdateCamera()
	CameraFollow(
		cc,
		sc.GetCenter().X,
		sc.GetCenter().Y,
	)

	rl.ClearBackground(rl.Black)

	rl.BeginMode2D(cc.Camera)

	DrawWorld(w, *cc)

	rl.DrawRectangleRec(sc.Rectangle, sc.Color)

	rl.EndMode2D()

	rl.EndDrawing()
}

func main() {
	rl.InitWindow(InitWindowWidth, InitWindowHeight, "Sekorathy")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	cameraController := NewCameraController()
	squareController := NewSquareController()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	sound := rl.LoadSound("sound.wav")
	defer rl.UnloadSound(sound)

	world := World{
		tiles: make(map[[2]int32]Tile),
	}

	buildMap(world, MovementLine)

	var recordedMovementLine string

	for !rl.WindowShouldClose() {
		handleInput(squareController, cameraController, sound, world)
		drawFrame(world, *squareController, cameraController)
		if getCurrentMovement() != 0 {
			recordedMovementLine += string(getCurrentMovement())
		}
		rl.DrawText(recordedMovementLine, 10, 10, 20, rl.Green)
	}
}
