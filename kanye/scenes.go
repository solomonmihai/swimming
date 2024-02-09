package kanye

import (
	"fmt"
)

type Scene interface {
	Update()
}

type BaseScene struct {
	entities []*Entity
}

func (scene *BaseScene) Entities() []*Entity {
	return scene.entities
}

type SceneManager struct {
	scenes map[string]Scene
	currentScene string
}

func NewSceneManager() SceneManager {
	return SceneManager{
		scenes: make(map[string]Scene),
	}
}

func (sm *SceneManager) AddScene(name string, scene Scene) bool {
	if _, found := sm.scenes[name]; found {
		return false
	}

	sm.scenes[name] = scene

	return true;
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
