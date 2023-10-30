# Golang Game Engine

This is a golang game enginer build ontop of the raylib
c bindings for golang. This engine does not feature any
UI, rather all interactions with this engine are done via
code. 

Submit bug reports and feature suggestions, or track changes in the
[Github Issue](https://github.com/jtheiss19/game-raylib/issues/new).

## Table of contents

- Requirements
- Basic Use - Hello World
    - Creating a custom component
    - Creating a custom object
    - Creating a custom system
- Troubleshooting
- FAQ
- Maintainers

## Requirements

This module uses the C bindings for raylib. Therefore you must have a
valid c compiler installed onto your machine for go to find and use. For
more detailed install instructions please see
[raylib-go](https://github.com/gen2brain/raylib-go#requirements)
on github for more detailed instructions

All other dependencies are managed by golang and one should have no issue
using them.

## Examples

Two Examples are found under the path `cmd/<example_name>/`

## Basic Use - Hello World

The main feature of this engine is the entity component system (ecs). In short a 
user creates a list of `components` that define a larger `object`. `Systems` then
loop over these `components` and find only the pairs of `components` they need,
then preform operations on just those `components`. This leaves `systems` blind
to `components` belonging to an `object` that it does not need. At larger scales 
this becomes signifficantly more efficient and easier to manage. Adding functionality
becomes as easy add defining a new `component` and appending it to old `objects`.

To start one must create a `world`. This will host all of your components. 
You may create many different `worlds` but for this example we will create just one.

``` go
func main() {
	// Create actual usable screen
	engine.SetupScreen()

	// Create world
	world := ecs.NewWorld()
	engine.BootstrapWorldRenderRaylib(world)
}
```

please note we are also adding some functionality to this world. Allthough not required
it is possible to create just a world and run it without any actual rendering system.
While this may be disired for servers, for this example we want a window. We call 
`engine.SetupScreen()` to create a screen for us and `engine.BootstrapWorldRenderRaylib(world)` to add some event hooks for drawing that are usefull. You could initialize these
yourself but these helper functions will preform the background tasks for us.

Now that we have a world lets populate it with some `systems`.

``` go
...
    // Create and add systems
	chunkRenderer := systems3d.NewChunkRenderingSystem()
	world.AddSystem(chunkRenderer)
	Renderer := systems3d.NewRenderingSystem()
	world.AddSystem(Renderer)
	pcs := systems3d.NewPlayerControllerSystem()
	world.AddSystem(pcs)
	modelManager := systems3d.NewModelLoadingSystem()
	world.AddSystem(modelManager)
...
```

Her we add a couple systems to the world using the `world.AddSystem(<Your System>)`. We will talk about these systems, how they are made, what they do, and how to make your own later, but for now this will get your application working with a few different systems. 

Now we need to populate the world with actual data. Something for these systems to itterate over.

``` go
...
	// Add objects to world
	world.AddEntity(objects3d.New3DPlayer(ecs.ID(uuid.New().String()), 0, 1, 0))
	world.AddEntity(objects3d.NewBlock3d(5, 1, -1, components3d.CRATE_TEX, 1))
	world.AddEntity(objects3d.NewBlock3d(5, 1, 0, components3d.CRATE_TEX, 1))
	world.AddEntity(objects3d.NewBlock3d(5, 1, 1, components3d.IMAGE_TEX, 9))
	world.AddEntity(objects3d.NewChunk(0, 0, 0))
...
```

This will add a vast array of different components. Each system will inspect each incomming
component to see if it can utilize it. If it can, it will cache it and run its functionality
over it next time it is called to run. 

Finally we have to start the systems, this is different from creating the window as it may 
already be on your screen. Here we are just simply kicking off a loop which will run over
each system and call their functionality. 

```
	// GameLoop
	engine.RunWorld(world)
```

### Creating a custom component

A little about components before moving on to creating them. Components hold only data,
not functionality, **ESPECIALLY NOT DATA MUTATING FUNCTIONALITY**. The ECS world is designed
with the principle that the data mutation will be happening only by systems. That being said...

Components are just any struct that matches an interface. We can add this interface by
using the struct inheritance below

``` go
type MyCustomComponent struct {
    *ecs.BaseComponent
    <Other Data You Want Here>
}
```

That's it. You now have a component that can register with the world and hold any data you
may want. 


### Creating a custom object

Objects are a logical abstraction. They aren't a defined thing in an ECS system. Rather they
are simply a collection of systems that share an identical key. Importing the `*ecs.BaseComponent` into your custom component will handle that functionality for you. 

Thus you are only responsible for grouping them togeather in the simplist way go can handle it. `[]ecs.Component`. 

**NOTE:** The components must be pointers to the underlying struct.

``` go
    myObject := []ecs.Component{
        &MyCustomComponent1,
        &MyCustomComponent2,
        &AEnginePredifinedComponent,
        ...
    }

    world.AddEntity(myObject)
```

This `AddEntity` function will handle registering your component with applicalable systems 
on the fly. So adding them to the world is as easy as grouping them. 

### Creating a custom system

Systems are where the magic happens. They are what acutally take stale lifeless data and 
begin to move it around. 

To make this happen you'll have to first define a system as below. 

``` go
    type MyCustomSystem struct {
        *ecs.BaseSystem
    }

    func NewMyCustomSystem() *MyCustomSystem {
        return &MyCustomSystem{
            BaseSystem:  &ecs.BaseSystem{},
        }
    }
```

This will give you an object that will allow itself to be managed by the world. It's important to have the helper function available, otherwise you may not initialize you system variables correctly. It also returns the system in the only form an instance of it should exist in, a pointer.

Next we need to tell the world what components to look for, and allow the system to find and use. 

``` go
// Comps
type RequiredMyCustomSystemComps struct {
	Grouping1 []*RequiredComponents1
    Grouping2 []*RequiredComponents2
}

type RequiredComponents1 struct {
	CustomComp1     *CustomComp1
	CustomComp2     *CustomComp2
}

type RequiredComponents2 struct {
	CustomComp3         *CustomComp3
	PredifinedComp1     *PredifinedComp1
	PredifinedComp2     *PredifinedComp2
}

func (ts *MyCustomSystem) GetRequiredComponents() interface{} {
	return &RequiredMyCustomSystemComps{
		Grouping1: []*RequiredComponents1{},
        Grouping2: []*RequiredComponents2{},
	}
}
```

This hook `GetRequiredComponents` is called by the world to findout what data types the system is looking for. These must be strictly difined data types (so don't use an interface). Above you can see the required components this system look for in an object can match either or both of the groupings `RequiredComponents1` or `RequiredComponents2`. To be considered matching requires the object to have a component of each specified value in `RequiredComponents1` or `RequiredComponents2` respectfully. 

Next once we have created and told the world what we want mounted to our system, we then need to add actual functionality. We do this as

``` go
func (ts *MyCustomSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredMyCustomSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}
}
```

Here `ts.TrackedEntities` contains all the components found that match our definitions previously. We cast it to the correct data type we know it must be (afterall we difined it earlier) and check for errors. If the cast fails, something internally must we wrong caused by either bad requirement definitions or passing by value instead of by reference. 

``` go
func (ts *MyCustomSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredMyCustomSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

    for _, obj := range entities.RequiredComponents1 {
        Do whatever functionality you want...
        Here Obj is 
        type RequiredComponents1 struct {
            CustomComp1     *CustomComp1
            CustomComp2     *CustomComp2
        }
    }

    for _, obj := range entities.RequiredComponents2 {
        Do whatever functionality you want...
        here Obj is
        type RequiredComponents2 struct {
            CustomComp3         *CustomComp3
            PredifinedComp1     *PredifinedComp1
            PredifinedComp2     *PredifinedComp2
        }
    }
}
```

Your functionality is completly left up to you. You may mutate the data in anyway you want. You may create new components and register them to the world or whatever you heart desires. The passed in var `dt` represents the time in milliseconds since the last update. At times this value may be equivlent to `0` to represent the world lagging, therefore it passes a 0 so noncritical or time based simulation event stop running for a call to allow the world machine to maintain 60 fps. This addhearence to respecting the `0` convention is completely up to your system.

Finally your system may want to preform an action, but only are mounting. This can be handled below 

``` go
func (ts *MyCustomSystem) Initilizer() {
    Init Code here...
}
```

Please note that all three of these functions are required to be present and defined on your custom system for it to able to used by the world. They do not need to be useful, or have actual content, but they must be defined. 

## Troubleshooting

To Be Written


## FAQ

**Q: Example Question.**

**A:** Example Answer


## Maintainers

- jthiss19 - [Github](https://github.com/jtheiss19)