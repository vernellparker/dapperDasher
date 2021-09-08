package main

import rl "github.com/gen2brain/raylib-go/raylib"

// window dimensions
const windowWidth int32 = 512
const windowHeight int32 = 380


var velocity int32 = 0


var scarfy rl.Texture2D
var scarfyRec rl.Rectangle 
var scarfyPos rl.Vector2
const gravity int32 = 1

var isInAir = false

const jumpVel int32 = -22

func init(){
	scarfy = rl.LoadTexture("../assets/textures/scarfy.png")

}

func main() {
	rl.InitWindow(windowWidth, windowHeight, "Dapper Dasher")
	rl.SetTargetFPS(60)
	scarfyRec.Width = float32(scarfy.Width)/6
	scarfyRec.Height = float32(scarfy.Height)
	scarfyRec.X = 0
	scarfyRec.Y = 0
	scarfyPos.X = float32(windowHeight)/2 - scarfyRec.Width/2
	scarfyPos.Y = float32(windowHeight)-scarfyRec.Height

	// Update
	for !rl.WindowShouldClose() {
		//start drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// ground check
		if int32(scarfyPos.Y) >= windowHeight-int32(scarfyRec.Height) {
			// is on ground
			velocity = 0
			isInAir = false
			scarfyPos.Y = float32(windowHeight - scarfy.Height)
		} else {
			// is in air
			isInAir = true
			velocity += gravity
		}

		// Jump Check
		if rl.IsKeyDown(rl.KeySpace) && !isInAir {
			velocity += jumpVel
		}

		scarfyPos.Y += float32(velocity)

		rl.DrawTextureRec(scarfy,scarfyRec,scarfyPos, rl.White)
		//stop drawing
		rl.EndDrawing()
	}
	rl.UnloadTexture(scarfy)
	rl.CloseWindow()
}
