package particles

import "project-particles/config"

func (s *System) GenerateParticles(num int) {
	for i := 0; i < num; i++ {
		p := createParticle()

		if s.Content.Len() >= config.General.MaxNumParticles {
			break
		}

		s.Content.PushFront(&p) // Add the particule address to the system
	}
}

func (s *System) GenerateParticlesWithCoord(num int, x float64, y float64) {
	for i := 0; i < num; i++ {
		p := createParticle()

		if s.Content.Len() >= config.General.MaxNumParticles {
			break
		}

		p.setPos(x, y)

		s.Content.PushFront(&p)
	}
}
