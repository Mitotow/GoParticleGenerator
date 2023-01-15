package particles

import (
	"project-particles/config"
)

func (p *Particle) update() {

	p.PositionY += p.SpeedY // Increase Y position of the particule with a certain speed
	p.PositionX += p.SpeedX // Same for the X position

	if config.General.EnableGravityEffect {
		p.gravity()
	}

	if config.General.EnableLifeSpan {
		if config.General.SmoothSuppression {
			if p.Life < -10 || p.ScaleX <= 0 || p.ScaleY <= 0 || p.Opacity <= 0 {
				p.Killed = true
			} else if p.Life < 1 {
				p.ScaleX, p.ScaleY, p.Opacity = p.ScaleX-0.1, p.ScaleY-0.1, p.Opacity-0.1
			}
		} else {
			if p.Life < 0 {
				p.Killed = true
			}
			if p.Life == 0 {
				p.ScaleX, p.ScaleY = 0, 0
			}
		}
	}

	if !p.InScreen() {
		// Particule is out of the screen
		if config.General.KillWhenOutOfScreen {
			p.Killed = true
			p.ScaleX, p.ScaleY = 0, 0
		} else {
			if config.General.RandomSpawn == true {
				// RandomSpawn enable then move the particule to random coordonates at the screen
				p.randomPos()
			} else {
				// RandomSpawn disable then move the particule to the choosen coordonates in config.json
				p.setSpawn()
			}
		}
	}

	p.Life--
}
