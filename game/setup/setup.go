package setup

import (
	"main/game"
	// "fmt"
	"main/game/scenes"
	"main/kanye"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

func Setup() {
	game.Game = kanye.NewGame("// swimming //")

	game.Game.Init(800, 600, 60)

	game.Game.Assets.LoadAssets([]kanye.AssetDesc{
		{
			Name:      "cube",
			Path:      "./assets/cube.obj",
			AssetType: kanye.MODEL,
		},
		{
			Name:      "test",
			Path:      "./assets/test.png",
			AssetType: kanye.TEXTURE,
		},
	})

	game.Game.Scenes.LoadScenes([]kanye.SceneDesc{
		{
			Name: "main",
			Scene: scenes.NewMainScene(),
		},
	})

	game.Game.Scenes.SetScene("main")

	game.Game.Start()

	// rl.SetTraceLogLevel(rl.LogError)
	// rl.InitWindow(800, 600, "hello")
	// rl.SetTargetFPS(60)

	// defer rl.CloseWindow()
	// // defer game.Assets.Close()
	// // defer game.Scenes.Close()

	// am := kanye.NewAssetsManager()

	// am.LoadAssets([]kanye.AssetDesc{
	// {
	// Name:      "test",
	// Path:      "./assets/test.png",
	// AssetType: kanye.IMAGE,
	// },
	// })

	// img := am.Image("test")
	// t := rl.LoadTextureFromImage(img)

	// sm := kanye.NewSceneManager()
	// sm.AddScene("main", scenes.NewMainScene(img))
	// sm.SetScene("main")

	// for !rl.WindowShouldClose() {
	// rl.ClearBackground(rl.SkyBlue)

	// rl.BeginDrawing()

	// rl.DrawTexture(t, 100, 100, rl.White)

	// // fmt.Println(t.Width)

	// sm.Update()

	// rl.EndDrawing()
}
