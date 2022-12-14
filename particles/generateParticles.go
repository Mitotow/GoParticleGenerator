package particles

import (
	"math/rand"
	"project-particles/config"
)

func (s System) GenerateParticles() {
	for i := 0; i < config.General.InitNumParticles; i++ {
		x := 0 + rand.Float64()*(float64(config.General.WindowSizeX)-0)
		y := 0 + rand.Float64()*(float64(config.General.WindowSizeY)-0)
		s.CreateParticle(x, y, 0+rand.Float64()*(1-0), 0.25, 1, 0, 0, 1, 0+rand.Float64()*(1-0))
	}
}
