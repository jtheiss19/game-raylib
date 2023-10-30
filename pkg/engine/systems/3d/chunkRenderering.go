package systems3d

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"
	"github.com/jtheiss19/game-raylib/pkg/engine/components"
	components3d "github.com/jtheiss19/game-raylib/pkg/engine/components/3d"

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
	TransformationComp *components3d.ChunkTransformation3DComponent
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
			TransformationComp: &components3d.ChunkTransformation3DComponent{},
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
			// Skip if model isn't loaded yet
			if chunkData.ModelComp.Model.Meshes == nil {
				continue
			}

			// Render Mesh
			transformMatricies := []rl.Matrix{}
			for _, chunkType := range chunkData.Chunk3DComp.Data {
				if chunkType.BlockType == 1 {
					scaleMatrix := rl.MatrixScale(
						chunkData.TransformationComp.Scale.X,
						chunkData.TransformationComp.Scale.Y,
						chunkData.TransformationComp.Scale.Z,
					)
					translateMatrix := rl.MatrixTranslate(
						chunkData.TransformationComp.Position.X+chunkType.RelativePosition.X,
						chunkData.TransformationComp.Position.Y+chunkType.RelativePosition.Y,
						chunkData.TransformationComp.Position.Z+chunkType.RelativePosition.Z,
					)
					transformMatrix := rl.MatrixMultiply(scaleMatrix, translateMatrix)
					transformMatricies = append(transformMatricies, transformMatrix)
				}
			}
			shader := chunkData.ModelComp.Model.Materials.Shader
			rl.SetShaderValue(shader, shader.GetLocation(rl.LocVectorView),
				[]float32{camera.Position.X, camera.Position.Y, camera.Position.Z}, rl.ShaderUniformVec3)
			rl.DrawMeshInstanced(
				*chunkData.ModelComp.Model.Meshes,
				*chunkData.ModelComp.Model.Materials,
				transformMatricies,
				len(transformMatricies),
			)
		}

		rl.EndMode3D()

		rl.DrawRectangle(8, 8, 80, 23, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(8, 8, 80, 23, rl.Blue)
		rl.DrawFPS(10, 10)
	}
}
