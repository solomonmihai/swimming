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
		{ "cube", "./assets/cube.obj", kanye.ASSET_MODEL, },
		{ "test", "./assets/test.png", kanye.ASSET_TEXTURE, },

		{ "tree", "./assets/models/tree/tree.obj", kanye.ASSET_MODEL, },
		{ "treeTexture", "./assets/models/tree/texture.png", kanye.ASSET_TEXTURE, },
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
