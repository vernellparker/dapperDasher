package main

import "github.com/gen2brain/raylib-go/raylib"

// window dimensions
const windowWidth int32 = 512
const windowHeight int32 = 380

const width int32 = 50
const height int32 = 80

var posY int32 = windowHeight - height
var velocity int32 = 0

func main() {
	rl.InitWindow(windowWidth, windowHeight,"Dapper Dasher")
	rl.SetTargetFPS(60)

	// Update
	for !rl.WindowShouldClose() {
		//start drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsKeyDown(rl.KeySpace){velocity += -10}

		posY += velocity

		rl.DrawRectangle(windowWidth/2, posY, width, height, rl.Blue)

		//stop drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
