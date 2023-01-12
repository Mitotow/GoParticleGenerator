package particles

import (
	"container/list"
	"project-particles/config"
)

func (s *System) particleSuppression(e *list.Element) {
	p, ok := e.Value.(*Particle)
	if !ok {
		return
	}

	if !config.General.SmoothSuppression {
		if p.Life < 0 {
			s.Content.Remove(e)

		}
		if p.Life == 0 {
			p.ScaleX, p.ScaleY = 0, 0
		}
	} else {
		if p.Life < -10 || p.ScaleX <= 0 || p.ScaleY <= 0 || p.Opacity <= 0 {
			s.Content.Remove(e)
		} else if p.Life < 1 {
			p.ScaleX, p.ScaleY, p.Opacity = p.ScaleX-0.1, p.ScaleY-0.1, p.Opacity-0.1
		}
	}
}
