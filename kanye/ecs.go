package kanye

import (
	"fmt"
	"reflect"
)

type Entity struct {
	id         int
	components map[ComponentId]Component
}

func (entity *Entity) Id() int {
	return entity.id
}

type ComponentId int
type Component interface{
	Id() ComponentId
}

type System interface {
	Update(scene *BaseScene)
}
type BaseSystem struct{}

var entityCount = 0

func EntityCount() int {
	return entityCount
}

func NewEntity() *Entity {
	entity := &Entity{
		id:         entityCount,
		components: make(map[ComponentId]Component),
	}
entityCount += 1

	return entity
}

func (entity *Entity) AddComponent(components ...Component) {
	for _, component := range components {
		if _, found := entity.components[component.Id()]; found {
			panic(fmt.Sprintf("[ecs] component id %d already added on entity %d", component.Id(), entity.Id()))
		}

		entity.components[component.Id()] = component
	}
}

func (entity *Entity) GetComponent(componentId ComponentId) *Component {
	component, found := entity.components[componentId]

	if !found {
		return nil
	}

	return &component
}

func (entity Entity) HasComponent(componentId ComponentId) bool {
	if entity.GetComponent(componentId) == nil {
		return false
	}

	return true
}

func (entity Entity) HasComponents(componentIds []ComponentId) bool {
	for _, componentId := range componentIds {
		if !entity.HasComponent(componentId) {
			return false
		}
	}

	return true
}

func (system *BaseSystem) IterateEntities(scene *BaseScene, f func(*Entity), requiredComponents ...ComponentId, ) {
	for _, ent := range scene.Entities() {
		if !ent.HasComponents(requiredComponents) {
			continue
		}

		f(ent)
	}
}

func name(t interface{}) string {
	name := reflect.TypeOf(t).String()

	if name[0] == '*' {
		return name[1:]
	}

	return name
}
