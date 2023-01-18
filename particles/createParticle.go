package particles

import (
	"project-particles/config"
)

var ParticuleNumber = 0

func createParticle() Particle {
	p := Particle{
		Id: ParticuleNumber,
		PositionX: float64(config.General.SpawnX), PositionY: float64(config.General.SpawnY),
		Rotation: config.General.Rotation,
		ScaleX:   config.General.ScaleX, ScaleY: config.General.ScaleX,
		ColorRed: config.General.ColorR, ColorGreen: config.General.ColorG, ColorBlue: config.General.ColorB,
		Opacity: config.General.Opacity,
		SpeedX:  config.General.SpeedX,
		SpeedY:  config.General.SpeedY,
		Life:    config.General.LifeSpan,
	}

	p.setSpawn()

	if config.General.RandomRotation {
		// RandomRotation enable
		p.randomRotation()
	}

	if config.General.RandomColor {
		// RandomColor enable
		p.randomColor() // Create a random color
		if config.General.BetterRandomColor {
			p.betterRandomColor()
		}
	}

	if config.General.RandomOpacity {
		// RandomOpacity enable
		p.randomOpacity()
	}

	if config.General.RandomScale {
		// RandomScale enable
		p.randomScale()
	}

	if config.General.RandomSpeed {
		p.randomSpeed()
	}

	if config.General.RandomLifeSpan {
		p.randomLifeSpan()
	}

	p.Width, p.Height = p.getSize()

	ParticuleNumber++

	return p
}
