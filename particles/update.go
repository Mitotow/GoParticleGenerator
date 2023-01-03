package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

var spawn = false
var particuleToAdd = float64(0)

func (s *System) Update() {

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(0)
	}

	particuleToAdd += config.General.SpawnRate - float64(int(config.General.SpawnRate))
	s.GenerateParticlesWithNumber(int(particuleToAdd + config.General.SpawnRate))

	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		p.PositionY += p.SpeedY
		p.PositionX += p.SpeedX
		if p.InScreen() == false {
			if config.General.RandomSpawn == true {
				p.PositionX, p.PositionY = randomPos()
			} else {
				p.PositionX, p.PositionY = float64(config.General.SpawnX), float64(config.General.SpawnY)
			}
			continue
		}
	}

	if config.General.SpawnRate == 0.5 && spawn == false {
		spawn = true
	}

}
