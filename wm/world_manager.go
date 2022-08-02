package wm

import "github.com/RoyalMCPE/world-manager/wm/world"

// path is the location to all the worlds
func NewManager(path string) world.WorldManager {
	return world.NewWorldManager(path)
}
