package scenes

import (
	"main/game/components"
	"main/game/entities"
	"main/game/systems"
	"main/kanye"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainScene struct {
  kanye.BaseScene
}

func NewMainScene() *MainScene {
  scene := &MainScene{}
  scene.Init()

  scene.AddSystem(&systems.ModelRenderer{})
  scene.AddSystem(systems.NewCameraController(&scene.BaseScene))

  ground := entities.NewGround(rl.Vector3Zero())
  groundTransform := (*ground.GetComponent(components.TRANSFORM_ID)).(*components.Transform)
  groundTransform.Scale = rl.NewVector3(100, 0.3, 100)

  scene.AddEntity(
    entities.NewPlayer(),
    ground,
  )

  rl.DisableCursor()

  return scene
}

func (scene *MainScene) Update() {}

func (scene *MainScene) Close() {}

