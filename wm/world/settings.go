package world

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

type Settings struct {
	SpawnPosition cube.Pos
	Dimension     world.Dimension
	Generator     world.Generator
}
