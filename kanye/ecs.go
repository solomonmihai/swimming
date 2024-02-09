package kanye

import "reflect"

type Entity struct {
	id int
	components map[string]Component
}

type Component interface{}

var entityCounter = 0

func NewEntity() *Entity {
	entity := &Entity{
		id: entityCounter,
		components: make(map[string]Component),
	}

	entityCounter += 1

	return entity
}

// TODO: add multiple components at once

func (entity *Entity) AddComponent(component Component) bool {
	componentName := name(component);

	if _, found := entity.components[componentName]; found {
		return false
	}

	entity.components[componentName] = component

	return true
}

func (entity *Entity) GetComponent(component Component) (*Component, bool) {
	component, found := entity.components[name(component)]

	if !found {
		return nil, false
	}

	return &component, true
}

func name(t interface{}) string {
	name := reflect.TypeOf(t).String()

	if name[0] == '*' {
		return name[1:]
	}

	return name
}
