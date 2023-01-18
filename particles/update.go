package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"project-particles/config"
)

var particuleToAdd float64 = 0
var tts = 0 // Temporisation counter (60 = 1s)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {

	if !config.General.SpawnOnMouse {
		if tts >= config.General.TimeToSpawn {
			// Every x frame generate new particle(s)

			if s.Content.Len() < config.General.MaxNumParticles || config.General.MaxNumParticles == 0 {
				// If the particle limit is not reached or the particule limit is unlimited
				particuleToAdd += config.General.SpawnRate - float64(int(config.General.SpawnRate))
				s.GenerateParticles(int(particuleToAdd) + int(config.General.SpawnRate)) // Generate an amount of particule at the screen
			}
			if particuleToAdd >= 1 {
				particuleToAdd = 0
			}

			tts = 0 // Restart the counter
		}
	} else if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		s.GenerateParticles(int(config.General.SpawnRate))
	}

	for e := s.Content.Front(); e != nil; {
		// Get each existing particules of the system
		p, ok := e.Value.(*Particle)
		next := e.Next()
		if !ok {
			continue
		}

		if p.Killed {
			s.Content.Remove(e)
		} else {
			if s.Collides(p) && config.General.EnableCollision {
				p.SpeedX *= -1
				p.SpeedY *= -1
			}
			p.update()
		}

		e = next
	}

	tts++ // Increase the temporisation counter
}
