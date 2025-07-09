package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	InitWindowWidth  = 800
	InitWindowHeight = 450
	TargetFPS = 240

	MovementLine = "/AAWDWWAAwwdDwWwaAaWwwDddSdDwWwwwDddSsssDdwwWaaAAAaaAAa"
)

func handleInput(sc *SquareController, cc *CameraController, w World) {
	sc.HandleInput(MovementLine)
	cc.HandleInput()
	placeTilesUsingCursor(w, *cc)
}

func drawFrame(w World, sc SquareController, cc *CameraController) {
	rl.BeginDrawing()

	cc.Update()
	cc.Follow(
		sc.GetCenter().X,
		sc.GetCenter().Y,
	)

	rl.ClearBackground(rl.Black)

	rl.BeginMode2D(cc.Camera)

	w.Draw(*cc)

	rl.DrawRectangleRec(sc.Rectangle, sc.Color)

	rl.EndMode2D()

	rl.EndDrawing()
}

func main() {
	rl.InitWindow(InitWindowWidth, InitWindowHeight, "Sekorathy")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(TargetFPS)

	cameraController := NewCameraController()
	squareController := NewSquareController()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	sound := rl.LoadSound("sound.wav")
	defer rl.UnloadSound(sound)

	world := World{
		tiles: make(map[[2]int32]Tile),
	}

	world.buildMap(MovementLine)

	for !rl.WindowShouldClose() {
		handleInput(squareController, cameraController, world)
		drawFrame(world, *squareController, cameraController)
	}
}
