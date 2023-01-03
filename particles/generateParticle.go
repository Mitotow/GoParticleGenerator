package particles

import (
	"math/rand"
	"project-particles/config"
)

func (s System) GenerateParticles() {
	content := s.Content
	x, y := float64(config.General.SpawnX), float64(config.General.SpawnY)
	for i := 0; i < config.General.InitNumParticles; i++ {
		if config.General.RandomSpawn {
			x, y = rand.Float64()*float64(config.General.WindowSizeX), rand.Float64()*float64(config.General.WindowSizeY)
		}
		spY := 1 + rand.Float64()*9
		spX := 1 + rand.Float64()*9
		p := createParticle(x, y, 1, 1, 1, 1, 1, 1, spY, spX)
		content.PushFront(&p)
	}
	s.Content = content
}
