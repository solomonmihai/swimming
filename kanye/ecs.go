package kanye

import (
	"reflect"
)

type Entity struct {
	id int
	components map[string]Component
}

func (entity *Entity) Id() int {
	return entity.id
}

type Component interface{}

type System interface {
	Update(scene *BaseScene)
}

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
// TODO: assign name to components, don't fetch them 
// by type name, sometimes we need to have multiple types of
// the same component on one entity (i.e. transforms)

func (entity *Entity) AddComponent(component Component) bool {
	componentName := name(component);

	if _, found := entity.components[componentName]; found {
		return false
	}

	entity.components[componentName] = component

	return true
}

func (entity *Entity) GetComponent(component Component) *Component {
	component, found := entity.components[name(component)]

	if !found {
		return nil
	}

	return &component
}

func GetComponent[T Component](entity *Entity) *T {
  comp, found := entity.components[name(*new(T))]

  if !found {
    return nil
  }

  return comp.(*T)
}

func name(t interface{}) string {
	name := reflect.TypeOf(t).String()

	if name[0] == '*' {
		return name[1:]
	}

	return name
}
