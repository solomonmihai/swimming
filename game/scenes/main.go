package scenes

import (
	"fmt"
	"main/kanye"
)

type MainScene struct {
  kanye.BaseScene

  counter int
}

func InitMainScene() *MainScene {
  scene := MainScene{}

  return &scene
}

func (scene *MainScene) Update() {
  fmt.Println(scene.counter)
}

func (scene *MainScene) Increment() {
  scene.counter += 1
}
