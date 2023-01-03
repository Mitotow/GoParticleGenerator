package particles

import "project-particles/config"

func (p Particle) InScreen() bool {
	return !(p.PositionX < float64(config.General.WindowSizeX) && p.PositionX > -10) || (p.PositionY < float64(config.General.WindowSizeY) && p.PositionY > -10)
}
