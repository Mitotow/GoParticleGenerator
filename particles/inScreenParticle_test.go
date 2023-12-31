package particles

import (
	"math/rand"
	"project-particles/config"
	"testing"
)

// This function is used to get the config.json file, because the main.go is not used when go test command is execute
func init() {
	config.Get("../config.json")
}

func TestInScreen1(t *testing.T) {
	p := Particle{
		PositionX: rand.Float64() * float64(config.General.WindowSizeX),
		PositionY: rand.Float64() * float64(config.General.WindowSizeY),
	}

	if p.InScreen() == false {
		t.Fail()
	}
}

func TestInScreen2(t *testing.T) {
	p := Particle{
		PositionX: rand.Float64() * float64(config.General.WindowSizeX),
		PositionY: rand.Float64() * float64(config.General.WindowSizeY),
	}

	if p.InScreen() == false {
		t.Fail()
	}
}

func TestOutScreen1(t *testing.T) {
	p := Particle{
		PositionX: float64(config.General.WindowSizeX) + 100,
		PositionY: float64(config.General.WindowSizeY) + 100,
	}

	if p.InScreen() == true {
		t.Fail()
	}
}

func TestOutScreen2(t *testing.T) {
	p := Particle{
		PositionX: -100,
		PositionY: -100,
	}

	if p.InScreen() == true {
		t.Fail()
	}
}
