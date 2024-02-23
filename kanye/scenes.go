package kanye

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene interface {
	Update()
	UpdateSystems()
	Close()
}

type BaseScene struct {
	entities []*Entity
	systems map[string]*System
	Camera *rl.Camera
}

func (scene *BaseScene) Init() {
	scene.systems = make(map[string]*System)	
}

func (scene *BaseScene) Entities() []*Entity {
	return scene.entities
}

func (scene *BaseScene) AddEntity(entities ...*Entity) {
	for _, ent := range entities {
		scene.entities = append(scene.entities, ent)
	}
}

func (scene *BaseScene) AddSystem(system System) bool {
	systemName := name(system);

	if _, found := scene.systems[systemName]; found {
		return false
	}

	scene.systems[systemName] = &system

	return true
}

func (scene *BaseScene) GetSystem(system System) *System {
	sys, found := scene.systems[name(system)]

	if !found {
		return nil
	}

	return sys
}

func (scene *BaseScene) UpdateSystems() {
	for _, system := range scene.systems {
		(*system).Update(scene)
	}
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

	scene.UpdateSystems()
	scene.Update()
}

func (sm *SceneManager) Close() {
	for _, scene := range(sm.scenes) {
		scene.Close()
	}
}
