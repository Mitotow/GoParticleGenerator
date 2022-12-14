package particles

import "project-particles/config"

func (s System) CreateParticle(x float64, y float64, rot float64, sX float64, sY float64, r float64, g float64, b float64, a float64) {
	s.Content.PushFront(&Particle{
		PositionX: x,
		PositionY: y,
		Rotation:  rot,
		ScaleX:    sX, ScaleY: sY,
		ColorRed: r, ColorGreen: g, ColorBlue: b,
		Opacity: a,
		life:    config.General.ParticleLifeSpan,
	})
}
