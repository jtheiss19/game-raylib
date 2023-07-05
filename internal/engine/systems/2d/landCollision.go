package systems2d

import (
	"fmt"
	"math"
	"rouge/internal/ecs"
	components2d "rouge/internal/engine/components/2d"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/sirupsen/logrus"
)

type LandCollisionSystem struct {
	*ecs.BaseSystem
}

func NewLandCollisionSystem() *LandCollisionSystem {
	return &LandCollisionSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredLandCollisionSystemComps struct {
	Land       []*RequireLand
	Collidable []*RequireCollidable
}

type RequireCollidable struct {
	Collision      *components2d.CollisionComponent
	Transformation *components2d.Transformation2DComponent
}

func (ts *LandCollisionSystem) GetRequiredComponents() interface{} {
	return &RequiredLandCollisionSystemComps{
		Land:       []*RequireLand{{}},
		Collidable: []*RequireCollidable{{}},
	}
}

// Functionality
func (ts *LandCollisionSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredLandCollisionSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	var tileWidth float32 = 10.0
	spriteSize := rl.Vector2{X: 10, Y: 10}

	for _, landEntity := range entities.Land {
		landPosition := landEntity.TransformComponent.Position

		landWidth := landEntity.LandComponent.Width
		landHeight := landEntity.LandComponent.Height
		for _, collidableEntity := range entities.Collidable {
			collidablePosisionTopLeft := collidableEntity.Transformation.Position
			collidablePosisionBottomRight := rl.Vector2Add(collidableEntity.Transformation.Position, spriteSize)

			if collidablePosisionTopLeft.X > landPosition.X && collidablePosisionTopLeft.X < landPosition.X+float32(landWidth)*tileWidth {
				if collidablePosisionTopLeft.Y > landPosition.Y && collidablePosisionTopLeft.Y < landPosition.Y+float32(landHeight)*tileWidth {

					// Top Left Sprite Corner in chunk cords
					localLandCordsTopLeft := rl.Vector2Subtract(collidablePosisionTopLeft, landPosition)
					landXCollisionTopLeft := math.Floor(float64(localLandCordsTopLeft.X / tileWidth))
					landYCollisionTopLeft := math.Floor(float64(localLandCordsTopLeft.Y / tileWidth))

					// Bottom Right Sprite Corner in chunk cords
					localLandCordsBottomRight := rl.Vector2Subtract(collidablePosisionBottomRight, landPosition)
					landXCollisionBottomRight := math.Floor(float64(localLandCordsBottomRight.X / tileWidth))
					landYCollisionBottomRight := math.Floor(float64(localLandCordsBottomRight.Y / tileWidth))

					collisions := []rl.Vector2{}
					for x := landXCollisionTopLeft; x <= landXCollisionBottomRight; x++ {
						for y := landYCollisionTopLeft; y <= landYCollisionBottomRight; y++ {
							// Make sure its in the chunk dataset
							if int(x) < landWidth && int(y) < landHeight {
								if landEntity.LandComponent.Data[int(y)*landWidth+int(x)] != 0 {
									collisions = append(collisions, rl.Vector2{X: float32(x), Y: float32(y)})
								}
							}
						}
					}

					// Draw Output
					rl.DrawRectangle(10, 40, 240, 20, rl.Fade(rl.SkyBlue, 0.5))
					rl.DrawRectangleLines(10, 40, 240, 20, rl.Blue)
					rl.DrawText(fmt.Sprint(collisions), 10, 40, 20, rl.Black)

					// Resolve Collisions
					// for _, collision := range collisions {
					// 	collidableEntity.Transformation.Position = rl.Vector2{0, 0}
					// }
				}
			}
		}
	}

}

func (ts *LandCollisionSystem) Initilizer() {
}
