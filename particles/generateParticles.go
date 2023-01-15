package particles

import "project-particles/config"

func (s System) GenerateParticles(num int) {
	for i := 0; i < num; i++ {
		var p Particle
		p = createParticle()

		if s.Content.Len() >= config.General.MaxNumParticles {
			return
		}

		s.Content.PushFront(&p) // Add the particule address to the system
	}
}
