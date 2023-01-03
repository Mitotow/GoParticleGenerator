package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"testing"
)

/** Test for the GenerateParticles function **/
// Check if the number of particules is equal to the InitNumParticles in the config.json
func Test_GenerateParticles_Number1(t *testing.T) {
	s := System{Content: list.New()}
	s.GenerateParticles()
	if s.Content.Len() != config.General.InitNumParticles {
		t.Fail()
	}
}

// Check if for a random number, the generation still work
func Test_GenerateParticles_Number2(t *testing.T) {
	s := System{Content: list.New()}
	x := rand.Intn(100)
	config.General.InitNumParticles = x
	s.GenerateParticles()
	if s.Content.Len() != x {
		t.Fail()
	}
}

/** Test for the GenerateParticlesWithNumber function **/
// Check if the number of particules is equal to the InitNumParticles in the config.json
func Test_GenerateParticlesWithNumber_Number1(t *testing.T) {
	s := System{Content: list.New()}
	s.GenerateParticlesWithNumber(5)
	if s.Content.Len() != 5 {
		t.Fail()
	}
}

// Check if for a random number, the generation still work
func Test_GenerateParticlesWithNumber_Number2(t *testing.T) {
	s := System{Content: list.New()}
	x := rand.Intn(100)
	s.GenerateParticlesWithNumber(x)
	if s.Content.Len() != x {
		t.Fail()
	}
}
