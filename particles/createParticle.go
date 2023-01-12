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
	var lifeSpan = config.General.LifeSpan                                            // Default particule life span

	if x == -1 {
		x = float64(config.General.WindowSizeX) / 2
	}

	if y == -1 {
		y = float64(config.General.WindowSizeY) / 2
	}

	if config.General.RandomSpawn {
		// RandomSpawn enable
		x, y = randomPos() // Create random spawn position
	}

	if config.General.RandomColor {
		// RandomColor enable
		r, g, b = rand.Float64(), rand.Float64(), rand.Float64() // Create a random color
		if config.General.BetterRandomColor {
			random := rand.Float64()
			if random <= 0.33 {
				r = rand.Float64() * (0.33)
				g = (0.33) + rand.Float64()*(0.33)
				b = (0.66) + rand.Float64()*(0.44)
			} else if random <= 0.66 {
				b = rand.Float64() * (0.33)
				r = (0.33) + rand.Float64()*(0.33)
				g = (0.66) + rand.Float64()*(0.44)
			} else {
				g = rand.Float64() * (0.33)
				b = (0.33) + rand.Float64()*(0.33)
				r = (0.66) + rand.Float64()*(0.44)
			}
		}
	}

	if config.General.RandomOpacity {
		// RandomOpacity enable
		o = config.General.MinOpacity + rand.Float64()*(config.General.MaxOpacity-config.General.MinOpacity)
	}

	if config.General.RandomScale {
		// RandomScale enable
		sX = config.General.MinScaleX + rand.Float64()*(config.General.MinScaleX-config.General.MinScaleX)
		sY = config.General.MinScaleY + rand.Float64()*(config.General.MinScaleY-config.General.MinScaleY)
	}

	if config.General.RandomSpeed {
		spX, spY = config.General.MinSpeedX+rand.Float64()*(config.General.MaxSpeedX-config.General.MinSpeedX), config.General.MinSpeedY+rand.Float64()*(config.General.MaxSpeedY-config.General.MinSpeedY)
		if rand.Float64() > 0.5 {
			spX *= -1
		}
		if rand.Float64() > 0.5 {
			spY *= -1
		}
	}

	if config.General.RandomLifeSpan {
		lifeSpan = config.General.MinLifeSpan + int(rand.Float64()*(float64(config.General.MaxLifeSpan-config.General.MinLifeSpan)))
	}

	return Particle{
		PositionX: x,
		PositionY: y,
		ScaleX:    sX, ScaleY: sY,
		ColorRed: r, ColorGreen: g, ColorBlue: b,
		Opacity: o,
		SpeedX:  spX,
		SpeedY:  spY,
		Life:    lifeSpan,
	}
}
