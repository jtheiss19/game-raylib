package example

import "github.com/jtheiss19/game-raylib/pkg/ecs"

type ExampleStruct struct {
	*ecs.BaseComponent
}

func NewExampleStruct() *ExampleStruct {
	return &ExampleStruct{
		BaseComponent: &ecs.BaseComponent{},
	}
}
