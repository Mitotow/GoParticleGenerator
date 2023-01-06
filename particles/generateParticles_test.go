package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"testing"
	"time"
)

// This function is used to get the config.json file, because the main.go is not used when go test command is execute
func init() {
	config.Get("../config.json")
}

func Test_GenerateParticles_Number1(t *testing.T) {
	s := System{Content: list.New()}
	s.GenerateParticles(config.General.InitNumParticles)
	if s.Content.Len() != config.General.InitNumParticles {
		t.Error("Il y a ", s.Content.Len(), " particules générées, il devrait y en avoir ", config.General.InitNumParticles, ".")
	}
}

func Test_GenerateParticles_Number2(t *testing.T) {
	s := System{Content: list.New()}
	rand.Seed(time.Now().UnixNano())
	var x = rand.Intn(100)
	config.General.InitNumParticles = x
	s.GenerateParticles(config.General.InitNumParticles)
	if s.Content.Len() != x {
		t.Error("Il y a ", s.Content.Len(), " particules générées, il devrait y en avoir ", x, ".")
	}
}

func Test_GenerateParticles_Number3(t *testing.T) {
	s := System{Content: list.New()}
	s.GenerateParticles(5)
	if s.Content.Len() != 5 {
		t.Error("Il y a ", s.Content.Len(), " particules générées, il devrait y en avoir 5.")
	}
}

func Test_GenerateParticles_Number4(t *testing.T) {
	s := System{Content: list.New()}
	rand.Seed(time.Now().UnixNano())
	var x = rand.Intn(100)
	s.GenerateParticles(x)
	if s.Content.Len() != x {
		t.Error("Il y a", s.Content.Len(), "particules générées, il devrait y en avoir", x, ".")
	}
}
