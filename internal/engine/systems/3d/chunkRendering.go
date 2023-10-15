package systems3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components3d "rouge/internal/engine/components/3d"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ChunkRenderingSystem struct {
	*ecs.BaseSystem
}

func NewChunkRenderingSystem() *ChunkRenderingSystem {
	return &ChunkRenderingSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredChunkRenderingSystemComps struct {
	Chunk  []*RequireChunkData
	Camera []*RequireChunkCameraData
}

type RequireChunkData struct {
	TransformationComp *components3d.GridTransformation3DComponent
	ModelComp          *components3d.Model3DComponent
	Chunk3DComp        *components3d.Chunk3DComponent
}

type RequireChunkCameraData struct {
	Player         *components.PlayerComponent
	Camera         *components3d.Camera3DComponent
	Transformation *components3d.Transformation3DComponent
}

func (ts *ChunkRenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredChunkRenderingSystemComps{
		Chunk: []*RequireChunkData{{
			TransformationComp: &components3d.GridTransformation3DComponent{},
			ModelComp:          &components3d.Model3DComponent{},
			Chunk3DComp:        &components3d.Chunk3DComponent{},
		}},
		Camera: []*RequireChunkCameraData{{
			Player:         &components.PlayerComponent{},
			Camera:         &components3d.Camera3DComponent{},
			Transformation: &components3d.Transformation3DComponent{},
		}},
	}
}

// Functionality
func (ts *ChunkRenderingSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredChunkRenderingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	if len(entities.Camera) > 0 {

		camera := entities.Camera[0].Camera.Camera

		// Draw
		rl.BeginMode3D(*camera)

		for _, chunkData := range entities.Chunk {
			for _, chunkType := range chunkData.Chunk3DComp.Data {
				if chunkType == 1 {
					rl.DrawModel(chunkData.ModelComp.Model, chunkData.TransformationComp.Position, chunkData.TransformationComp.Scale.X, rl.White)
				}
			}
		}

		rl.DrawGrid(100, 2)

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 220, 70, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 220, 70, rl.Blue)

		rl.DrawFPS(10, 10)

	}
}
