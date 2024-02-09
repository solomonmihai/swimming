package main

import (
	"main/game/components"
	"main/game/scenes"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SquareRenderSystem struct{}

func (sqs *SquareRenderSystem) Update(entities []*kanye.Entity) {
	for _, e := range entities {

		transformPtr, _ := e.GetComponent(components.Transform{})
		transform := (*transformPtr).(*components.Transform)

		srcPtr, _ := e.GetComponent(components.SquareRender{})
		src := (*srcPtr).(*components.SquareRender)

		transform.Pos.X += 1

		rl.DrawRectangle(int32(transform.Pos.X), int32(transform.Pos.Y), int32(transform.Scale.X), int32(transform.Scale.Y), src.Color)
	}
}

func main() {
	// rl.SetConfigFlags(rl.FlagWindowUndecorate)
	rl.InitWindow(800, 600, "// swimming //")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetTraceLogLevel(rl.LogError)

	sceneManager := kanye.NewSceneManager()

	scene := scenes.InitMainScene()

	sceneManager.AddScene("main", scene)

  sceneManager.SetScene("main")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)

    sceneManager.Update()

    scene.Increment()

		rl.EndDrawing()
	}
}
