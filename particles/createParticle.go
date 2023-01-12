package particles

import (
	"math/rand"
	"project-particles/config"
)

func createParticle() Particle {
	var x, y = float64(config.General.SpawnX), float64(config.General.SpawnY)         // Default particle spawn position
	var r, g, b = config.General.ColorR, config.General.ColorG, config.General.ColorB // Default particle color
	var o = config.General.Opacity                                                    // Default particule opacity
	var sX, sY = config.General.ScaleX, config.General.ScaleY                         // Default particule scale
	var spX, spY = config.General.SpeedX, config.General.SpeedY                       // Default particule speed

	if config.General.RandomSpawn {
		// RandomSpawn enable
		x, y = randomPos() // Create random spawn position
	}

	if config.General.RandomColor {
		// RandomColor enable
		r, g, b = rand.Float64(), rand.Float64(), rand.Float64() // Create a random color
	}

	if config.General.RandomOpacity {
		// RandomOpacity enable
		o = rand.Float64()
	}

	if config.General.RandomScale {
		// RandomScale enable
		sX, sY = rand.Float64(), rand.Float64()
	}

	if config.General.RandomSpeed {
		spX, spY = 1+rand.Float64()*(config.General.MaxSpeedX-1), 1+rand.Float64()*(config.General.MaxSpeedY-1)
		if rand.Float64() > 0.5 {
			spX *= -1
		}
		if rand.Float64() > 0.5 {
			spY *= -1
		}
	}

	return Particle{
		PositionX: x,
		PositionY: y,
		ScaleX:    sX, ScaleY: sY,
		ColorRed: r, ColorGreen: g, ColorBlue: b,
		Opacity: o,
		SpeedX:  spX,
		SpeedY:  spY,
		Life:    config.General.LifeSpan,
	}
}
