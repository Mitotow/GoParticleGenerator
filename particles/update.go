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

func (s *System) Update() {

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(0)
	}

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
}
