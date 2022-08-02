package world

import (
	"fmt"
	"os"

	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/mcdb"
	"github.com/df-mc/goleveldb/leveldb/opt"
)

func NewWorldManager(path string) WorldManager {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	return WorldManager{path: path, worlds: make(map[string]*World)}
}

type WorldManager struct {
	path   string
	worlds map[string]*World
}

func (wm *WorldManager) Create(name string, settings *Settings) (*World, error) {
	provider, err := mcdb.New(fmt.Sprintf("%s/%s", wm.path, name), opt.DefaultCompression)
	if err != nil {
		return nil, err
	}

	provider.SaveSettings(&world.Settings{
		Name:  name,
		Spawn: settings.SpawnPosition,
	})

	config := world.Config{
		Provider:  provider,
		Generator: settings.Generator,
		Dim:       settings.Dimension,
	}

	w := &World{world: config.New()}
	wm.worlds[name] = w
	return w, nil
}

func (wm WorldManager) World(name string) *World {
	w, ok := wm.worlds[name]
	if !ok {
		return nil
	}
	return w
}

func (wm *WorldManager) Load(name string) (*World, error) {
	path := fmt.Sprintf("%s/%s", wm.path, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("The world at %v does not exist", path)
	}

	pro, err := mcdb.New(path, opt.DefaultCompression)
	if err != nil {
		return nil, err
	}

	config := world.Config{
		Provider: pro,
	}

	w := &World{world: config.New()}
	wm.worlds[name] = w
	return w, nil
}

func (wm *WorldManager) Unload(w *World) {
	delete(wm.worlds, w.world.Name())
	w.Close()
}
