package setup

import (
	"main/game"
	"main/game/scenes"
	"main/kanye"
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
}
