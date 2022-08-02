package world

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type World struct {
	world *world.World
}

func (w World) TransferTo(p *player.Player) {
	p.World().RemoveEntity(p)
	w.world.AddEntity(p)
}

func (w World) Close() {
	w.world.Close()
}
