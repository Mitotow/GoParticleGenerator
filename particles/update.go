package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
	"project-particles/config"
)

var particuleToAdd = float64(0)

var tts = 0 // Temporisation counter (60 = 1s)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		// Exit the process when Escape key is released
		os.Exit(0)
	}

	if tts >= config.General.SpawnFrameRate {
		// Every x frame generate new particle(s)
		particuleToAdd += config.General.SpawnRate - float64(int(config.General.SpawnRate))
		s.GenerateParticlesWithNumber(int(particuleToAdd) + int(config.General.SpawnRate)) // Generate an amount of particule at the screen

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

		if p.InScreen() == false {
			// Particule is out of the screen
			if config.General.RandomSpawn == true {
				// RandomSpawn enable then move the particule to random coordonates at the screen
				p.PositionX, p.PositionY = randomPos()
			} else {
				// RandomSpawn disable then move the particule to the choosen coordonates in config.json
				p.PositionX, p.PositionY = float64(config.General.SpawnX), float64(config.General.SpawnY)
			}
		}
	}

	tts++ // Increase the temporisation counter

}
