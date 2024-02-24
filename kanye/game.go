package kanye

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Scenes *SceneManager
	Assets *AssetsManager
	title string
	appendFpsToTitle bool
}

func NewGame(title string) *Game {
	return &Game{
		Scenes: NewSceneManager(),
		Assets: NewAssetsManager(),
		title: title,
		appendFpsToTitle: true,
	}
}

func (game *Game) Init(width int, height int, fps int) {
	rl.SetTraceLogLevel(rl.LogError)
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(width), int32(height), game.title)
}

func (game *Game) Start() {
	defer func() {
		game.Scenes.Close()
		game.Assets.Close()
		rl.CloseWindow()
	}()

	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.SkyBlue)

		rl.BeginDrawing()
		game.Scenes.Update()
		rl.EndDrawing()

		if game.appendFpsToTitle {
			rl.SetWindowTitle(fmt.Sprintf("%s [%d fps]", game.title, rl.GetFPS()))
		}
	}
}
