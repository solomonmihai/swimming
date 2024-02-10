package scenes

import (
	// "fmt"
	"fmt"
	"main/game"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainScene struct {
  kanye.BaseScene

  camera rl.Camera
  model rl.Model
  texture rl.Texture2D
}

var t rl.Texture2D

func NewMainScene() *MainScene {
  scene := &MainScene{}

  fmt.Println("init main scene")

  // scene.camera = rl.NewCamera3D(rl.NewVector3(50, 50, 50), rl.NewVector3(0, 10, 0), rl.NewVector3(0, 1, 0), 45, rl.CameraPerspective)

  // scene.model = game.Game.Assets.Model("cube")

  // img := rl.LoadImage("./assets/test.png")
  // img := game.Game.Assets.Image("test")
  // t = game.Game.Assets.Texture("test")

  // game.Game.Assets.Image("test")
  // game.Game.Assets.Model("cube")

  // fmt.Println(img.Width)
  // fmt.Println(scene.texture.Width)

  scene.texture = game.Game.Assets.Texture("test")

  return scene
}

func (scene *MainScene) Update() {

  // fmt.Println("ce pula mea")

  // fmt.Println(scene.texture.Width)

  rl.DrawTexture(scene.texture, 100, 100, rl.White)

  // rl.BeginMode3D(scene.camera)

  // rl.DrawModel(scene.model, rl.Vector3Zero(), 3.0, rl.Beige)

  // rl.EndMode3D()
}

func (scene *MainScene) Close() { 
  // rl.UnloadTexture(scene.texture)
}

