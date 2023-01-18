package particles

func (s *System) Collides(p *Particle) bool {
	for e := s.Content.Front(); e != nil; e = e.Next() {
		other := e.Value.(*Particle)
		if other.Id == p.Id {
			continue
		}

		if (p.PositionY >= other.PositionY+other.Height && other.PositionY+other.Height <= p.PositionY+p.Height) ||
			(p.PositionY+p.Height >= other.PositionY && other.PositionY <= p.PositionY) {
			if (p.PositionX+p.Width <= other.PositionX+other.Width && other.PositionX+other.Width <= p.PositionX+p.Width) ||
				(other.PositionX <= p.PositionX+p.Width && other.PositionX+other.Width >= p.PositionX) {
				return true
			}
		}
	}
	return false
}
