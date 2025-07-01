package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	InitWindowWidth  = 800
	InitWindowHeight = 450
)

func main() {
	rl.InitWindow(InitWindowWidth, InitWindowHeight, "Sekorathy")
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}
}
