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

var entityCounter = 0

func NewEntity() *Entity {
	entity := &Entity{
		id:         entityCounter,
		components: make(map[ComponentId]Component),
	}

	entityCounter += 1

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

func (entity Entity) HasComponents(componentIds ...ComponentId) bool {
	for _, componentId := range componentIds {
		if entity.GetComponent(ComponentId(componentId)) == nil {
			return false
		}
	}

	return true
}

func name(t interface{}) string {
	name := reflect.TypeOf(t).String()

	if name[0] == '*' {
		return name[1:]
	}

	return name
}
