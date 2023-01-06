package particles

func createParticle(x, y, sX, sY, r, g, b, o, spX, spY float64) Particle {

	return Particle{
		PositionX: x,
		PositionY: y,
		ScaleX:    sX, ScaleY: sY,
		ColorRed: r, ColorGreen: g, ColorBlue: b,
		Opacity: o,
		SpeedX:  spX,
		SpeedY:  spY,
	}
}
