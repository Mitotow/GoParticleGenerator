package particles

import (
	"math/rand"
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	s.GenerateParticles()

	if config.General.RainingMode == true {
		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)

			if p.InScreen() == false {
				p.ScaleX, p.ScaleY = 0, 0
				s.Content.Remove(e)
				continue
			}

			if p.life <= 0 {
				p.ScaleX, p.ScaleY = 0, 0
				s.Content.Remove(e)
				continue
			} else {
				p.life--
			}

			p.PositionY += 1 + rand.Float64()*(10-1)
			p.PositionX += 0 + rand.Float64()*(3-0)
		}
	}
}
