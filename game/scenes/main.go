package scenes

import (
	"main/game"
	"main/game/components"
	"main/game/systems"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainScene struct {
  kanye.BaseScene

  model rl.Model
  texture rl.Texture2D
}

var t rl.Texture2D

func NewMainScene() *MainScene {
  scene := &MainScene{}
  scene.Init()

  scene.AddSystem(&systems.ModelRenderer{})

  player := kanye.NewEntity()

  player.AddComponent(components.NewModel(game.Game.Assets.Model("cube")))
  player.AddComponent(components.DefaultTransform())

  scene.AddEntity(player)

  scene.Camera = rl.NewCamera3D(rl.NewVector3(50, 50, 50), rl.NewVector3(0, 10, 0), rl.NewVector3(0, 1, 0), 45, rl.CameraPerspective)

  rl.DisableCursor()

  return scene
}

func (scene *MainScene) Update() {
  rl.UpdateCamera(&scene.Camera, rl.CameraFirstPerson)
}

func (scene *MainScene) Close() { 
}

