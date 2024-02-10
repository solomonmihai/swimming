package kanye

import (
	"fmt"
)

type Scene interface {
	Update()
	Close()
}

type BaseScene struct {
	entities []*Entity
}

func Entities(scene *BaseScene) []*Entity {
	return scene.entities
}

func (scene *BaseScene) Entities() []*Entity {
	return scene.entities
}

type SceneDesc struct {
	Name string
	Scene Scene
}

type SceneManager struct {
	scenes map[string]Scene
	currentScene string
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes: make(map[string]Scene),
	}
}

func (sm *SceneManager) LoadScenes(scenesDesc []SceneDesc) {
	for _, desc := range scenesDesc {
		if _, found := sm.scenes[desc.Name]; found {
			panic(fmt.Sprintf("[scenes] scene %s already added", desc.Name))
		}

		sm.scenes[desc.Name] = desc.Scene
	}
}

func (sm *SceneManager) SetScene(newScene string) bool {
	if _, found := sm.scenes[newScene]; !found {
		fmt.Println("scene not found")
		return false;
	}

	sm.currentScene = newScene

	return true
}

func (sm *SceneManager) Update() {
	if sm.currentScene == "" {
		return
	}

	scene, found := sm.scenes[sm.currentScene]

	if !found {
		return
	}

	scene.Update()
}

func (sm *SceneManager) Close() {
	for _, scene := range(sm.scenes) {
		scene.Close()
	}
}
