package particles

import "project-particles/config"

func (p Particle) InScreen() bool {
	if p.PositionX > float64(config.General.WindowSizeX+10) || p.PositionY > float64(config.General.WindowSizeY+10) {
		return false
	}
	return true
}
