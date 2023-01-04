package particles

func (p Particle) InScreen() bool {
	wX, wY := getWindowSize()                                                                                         // Get
	return (-10 < p.PositionX && p.PositionX < float64(wX)+10) || (-10 < p.PositionY && p.PositionY < float64(wY)+10) // return a condition that return true or false to know if the particule is inside the window
}
