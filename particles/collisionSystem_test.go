package particles

import (
	"container/list"
	"testing"
)

func Test_Collision(t *testing.T) {
	s := System{
		Content: list.New(),
	}

	s.Content.PushFront(&Particle{
		Id:        1,
		PositionX: 100,
		PositionY: 100,
		Width:     10,
		Height:    10,
	})

	s.Content.PushFront(&Particle{
		Id:        2,
		PositionX: 100,
		PositionY: 110,
		Width:     10,
		Height:    10,
	})

	if !s.Collides(s.Content.Front().Value.(*Particle)) {
		t.Error("Il y a une collision entre p1 et p2, mais Collides() renvois False")
	}
}
