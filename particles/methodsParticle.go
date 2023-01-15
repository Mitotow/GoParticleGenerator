package particles

import (
	"math/rand"
	"project-particles/config"
)

// Color
func (p *Particle) randomColor() {
	p.ColorRed, p.ColorGreen, p.ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
}

func (p *Particle) betterRandomColor() {
	random := rand.Float64()
	var r, g, b float64
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
	p.ColorRed, p.ColorGreen, p.ColorBlue = r, g, b
}

// LifeSpan
func (p *Particle) randomLifeSpan() {
	p.Life = config.General.MinLifeSpan + int(rand.Float64()*(float64(config.General.MaxLifeSpan-config.General.MinLifeSpan)))
}

// Opacity
func (p *Particle) randomOpacity() {
	p.Opacity = config.General.MinOpacity + rand.Float64()*(config.General.MaxOpacity-config.General.MinOpacity)
}

// Spawn / Position
func (p *Particle) randomPos() (float64, float64) {
	return rand.Float64() * float64(config.General.WindowSizeX), rand.Float64() * float64(config.General.WindowSizeY)
}

func (p *Particle) returnSpawn() {
	p.PositionX, p.PositionY = float64(config.General.SpawnX), float64(config.General.SpawnY)
	if p.PositionX == -1 {
		p.centerX()
	}
	if p.PositionY == -1 {
		p.centerY()
	}
}

func (p *Particle) centerX() {
	p.PositionX = float64(config.General.WindowSizeX) / 2
}

func (p *Particle) centerY() {
	p.PositionY = float64(config.General.WindowSizeY) / 2
}

// Rotation
func (p *Particle) randomRotation() {
	p.Rotation = config.General.MinRotation + rand.Float64()*(config.General.MaxRotation-config.General.MinRotation)
}

// Scale
func (p *Particle) randomScale() {
	p.ScaleX = config.General.MinScaleX + rand.Float64()*(config.General.MaxScaleX-config.General.MinScaleX)
	p.ScaleY = config.General.MinScaleY + rand.Float64()*(config.General.MaxScaleY-config.General.MinScaleY)
}

// Speed
func (p *Particle) randomSpeed() {
	p.SpeedX, p.SpeedY = config.General.MinSpeedX+rand.Float64()*(config.General.MaxSpeedX-config.General.MinSpeedX), config.General.MinSpeedY+rand.Float64()*(config.General.MaxSpeedY-config.General.MinSpeedY)
}
