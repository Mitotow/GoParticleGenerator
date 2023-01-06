package particles

import (
	"math/rand"
	"project-particles/config"
	"testing"
	"time"
)

// This function is used to get the config.json file, because the main.go is not used when go test command is execute
func init() {
	config.Get("../config.json")
}

func Test_GenerateParticles_RandomPosition(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	p1 := createParticle()
	p2 := createParticle()

	if p1 == p2 {
		t.Fail()
	}
}
