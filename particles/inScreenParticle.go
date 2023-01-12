package particles

import (
	"project-particles/assets"
	"project-particles/config"
)

func (p Particle) InScreen() bool {
	// wX and wY global value
	var wX, wY = config.General.WindowSizeX, config.General.WindowSizeY
	return ((float64(assets.PImageSizeX)*p.ScaleX*-1)-10 < p.PositionX && p.PositionX < float64(wX)+(float64(assets.PImageSizeX)*p.ScaleX+10)) || ((float64(assets.PImageSizeY)*p.ScaleX*-1)-10 < p.PositionY && p.PositionY < float64(wY)+(float64(assets.PImageSizeY)*p.ScaleX+10)) // return a condition that return true or false to know if the particule is inside the window
}
