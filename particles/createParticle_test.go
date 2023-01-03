package tests

import (
	"math/rand"
	"project-particles/config"
	"project-particles/particles"
	"testing"
)

func TestSame1(t *testing.T) {
	p1 := particles.Particle{
		PositionX: 1,
		PositionY: 1,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
		SpeedX:  1,
		SpeedY:  1,
	}

	p2 := particles.createParticle(1, 1, 1, 1, 1, 1, 1, 1, 1, 1)

	if p1 != p2 {
		t.Fail()
	}
}

func TestSame2(t *testing.T) {
	p1 := particles.Particle{
		PositionX: rand.Float64() * float64(config.General.WindowSizeX),
		PositionY: rand.Float64() * float64(config.General.WindowSizeX),
		ScaleX:    rand.Float64(), ScaleY: rand.Float64(),
		ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
		Opacity: rand.Float64(),
		SpeedX:  rand.Float64() * 10,
		SpeedY:  rand.Float64() * 10,
	}

	p2 := particles.createParticle(p1.PositionX, p1.PositionY, p1.ScaleX, p1.ScaleY, p1.ColorRed, p1.ColorGreen, p1.ColorBlue, p1.Opacity, p1.SpeedX, p1.SpeedY)

	if p1 != p2 {
		t.Fail()
	}
}
