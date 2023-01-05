package particles

import (
	"math/rand"
	"project-particles/config"
)

func (s System) GenerateParticles(num int) {
	var x, y = float64(config.General.SpawnX), float64(config.General.SpawnY) // Default spawn configurate in config.json
	var r, g, b = config.General.ColorR, config.General.ColorG, config.General.ColorB
	content := s.Content
	for i := 0; i < num; i++ {
		if config.General.RandomSpawn {
			// RandomSpawn enable
			x, y = randomPos() // Get random coordonates
		}
		if config.General.RandomColor {
			// RandomColor enable
			r, g, b = rand.Float64(), rand.Float64(), rand.Float64() // Create a random color
		}
		spY := rand.Float64() * 10                            // Create random speed in Y
		spX := rand.Float64() * 10                            // Create random speed in X
		p := createParticle(x, y, 1, 1, r, g, b, 1, spY, spX) // Create the particule
		content.PushFront(&p)                                 // Add the particule address to the system
	}
	s.Content = content
}
