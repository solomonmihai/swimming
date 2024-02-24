package scenes

import (
	"main/game"
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

  scene.AddSystem(
    &systems.PlayerController{},
    systems.NewCameraController(&scene.BaseScene),

    &systems.ModelRenderer{},
    systems.NewDebugSystem(),
  )

  groundTransform := components.DefaultTransform()
  groundTransform.Scale = rl.NewVector3(50, 0.3, 50)
  groundModel := components.NewModel(game.Game.Assets.Model("cube"))
  groundModel.Color = &rl.Purple
  ground := entities.NewBox(groundTransform, groundModel)

  scene.AddEntity(
    entities.NewPlayer(),
    ground,
  )

  trees := []*kanye.Entity{}
  for i := 0; i < 30; i += 1 {
    treeTransform := components.DefaultTransform()
    scale := float32(rl.GetRandomValue(300, 800)) / 100.0
    treeTransform.Scale = rl.NewVector3(scale, scale, scale)
    treeTransform.Pos = rl.NewVector3(
      float32(rl.GetRandomValue(-50, 50)),
      scale,
      float32(rl.GetRandomValue(-50, 50)),
    )
    treeTransform.Rotation = float32(rl.GetRandomValue(0, 360))
    treeModel := components.NewModel(game.Game.Assets.Model("tree"))
    treeModel.SetTexture(game.Game.Assets.Texture("treeTexture"))
    tree := entities.NewBox(treeTransform, treeModel)

    trees = append(trees, tree)
  }

  scene.AddEntities(trees)

  return scene
}

func (scene *MainScene) Update() {}
func (scene *MainScene) Close() {}

