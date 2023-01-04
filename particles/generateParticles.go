package particles

import (
	"math/rand"
	"project-particles/config"
)

func (s System) GenerateParticles() {
	content := s.Content
	x, y := float64(config.General.SpawnX), float64(config.General.SpawnY) // Default spawn configurate in config.json
	for i := 0; i < config.General.InitNumParticles; i++ {
		// InitNumParticles is the amount of particule(s) to create
		if config.General.RandomSpawn {
			// RandomSpawn enable
			x, y = randomPos() // Get random coordonates
		}
		spY := 1 + rand.Float64()*9                           // Create random speed in Y
		spX := 1 + rand.Float64()*9                           // Create random speed in X
		p := createParticle(x, y, 1, 1, 1, 1, 1, 1, spY, spX) // Create the particule
		content.PushFront(&p)                                 // Add the particule address to the system
	}
	s.Content = content
}

func (s System) GenerateParticlesWithNumber(num int) {
	content := s.Content
	x, y := float64(config.General.SpawnX), float64(config.General.SpawnY) // Default spawn configurate in config.json
	for i := 0; i < num; i++ {
		if config.General.RandomSpawn {
			// RandomSpawn enable
			x, y = randomPos() // Get random coordonates
		}
		spY := 1 + rand.Float64()*9                           // Create random speed in Y
		spX := 1 + rand.Float64()*9                           // Create random speed in X
		p := createParticle(x, y, 1, 1, 1, 1, 1, 1, spY, spX) // Create the particule
		content.PushFront(&p)                                 // Add the particule address to the system
	}
	s.Content = content
}
