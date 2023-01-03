package tests

import (
	"math/rand"
	"project-particles/config"
	"project-particles/particles"
	"testing"
)

func testInScreen1(t *testing.T) {
	p := particles.Particle{
		PositionX: rand.Float64() * float64(config.General.WindowSizeX),
		PositionY: rand.Float64() * float64(config.General.WindowSizeY),
	}

	if p.InScreen() == true {
		t.Fail()
	}
}

func testInScreen2(t *testing.T) {
	p := particles.Particle{
		PositionX: rand.Float64() * float64(config.General.WindowSizeX),
		PositionY: rand.Float64() * float64(config.General.WindowSizeY),
	}

	if p.InScreen() == true {
		t.Fail()
	}
}

func testOutScreen1(t *testing.T) {
	p := particles.Particle{
		PositionX: float64(config.General.WindowSizeX + 100),
		PositionY: float64(config.General.WindowSizeY + 100),
	}

	if p.InScreen() == false {
		t.Fail()
	}
}

func testOutScreen2(t *testing.T) {
	p := particles.Particle{
		PositionX: float64(config.General.WindowSizeX + 100),
		PositionY: -100,
	}

	if p.InScreen() == false {
		t.Fail()
	}
}
