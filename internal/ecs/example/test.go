package example

import "rouge/internal/ecs"

type ExampleStruct struct {
	*ecs.BaseComponent
}

func NewExampleStruct() *ExampleStruct {
	return &ExampleStruct{
		BaseComponent: &ecs.BaseComponent{},
	}
}
