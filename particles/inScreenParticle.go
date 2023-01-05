package particles

import "project-particles/config"

func (p Particle) InScreen() bool {
	// wX and wY global value
	var wX, wY = config.General.WindowSizeX, config.General.WindowSizeY
	return (-10 < p.PositionX && p.PositionX < float64(wX)+10) || (-10 < p.PositionY && p.PositionY < float64(wY)+10) // return a condition that return true or false to know if the particule is inside the window
}
