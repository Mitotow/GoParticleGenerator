package particles

func (s System) GenerateParticles(num int) {
	for i := 0; i < num; i++ {
		p := createParticle()   // Create the particule
		s.Content.PushFront(&p) // Add the particule address to the system
	}
}
