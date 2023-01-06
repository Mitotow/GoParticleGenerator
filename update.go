package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {

	g.system.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// Exit the process when Escape key is released
		os.Exit(0)
	}

	return nil
}
