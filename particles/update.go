package particles

import (
	"fmt"
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

	if tts >= config.General.TimeToSawn {
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

	for e := s.Content.Front(); e != nil; e = e.Next() {
		// Get each existing particules of the system
		p := e.Value.(*Particle)

		p.PositionY += p.SpeedY // Increase Y position of the particule with a certain speed
		p.PositionX += p.SpeedX // Same for the X position

		fmt.Println("===================\nPosX :", p.PositionX, "\nPosY :", p.PositionY)

		if p.Killed {
			s.Content.Remove(e)
		}

		if config.General.EnableLifeSpan {
			go s.particleSuppression(e) // Utilisation of goroutine for optimisation
		}

		if !p.InScreen() {
			// Particule is out of the screen
			if config.General.KillWhenOutOfScreen {
				p.Killed = true
				p.ScaleX, p.ScaleY = 0, 0
				continue
			}
			if config.General.RandomSpawn == true {
				// RandomSpawn enable then move the particule to random coordonates at the screen
				p.PositionX, p.PositionY = randomPos()
			} else {
				// RandomSpawn disable then move the particule to the choosen coordonates in config.json
				p.PositionX, p.PositionY = float64(config.General.SpawnX), float64(config.General.SpawnY)
			}
		}

		p.Life--
	}

	tts++ // Increase the temporisation counter
}
