package main

import rl "github.com/gen2brain/raylib-go/raylib"

const FPS int32 = 30

const WIDTH int32 = 800
const HEIGHT int32 = 800

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "multiSandboxGame")

	rl.SetTargetFPS(FPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(300, 300, 20, 20, rl.SkyBlue)

		rl.EndDrawing()
	}
}
