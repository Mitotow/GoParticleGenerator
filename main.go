package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
)

// main est la fonction principale du projet. Elle commence par lire le fichier
// de configuration, puis elle charge en mémoire l'image d'une particule. Elle
// initialise ensuite la fenêtre d'affichage, puis elle crée un système de
// particules encapsulé dans un "game" et appelle la fonction RunGame qui se
// charge de faire les mise-à-jour (Update) et affichages (Draw) de manière
// régulière.
func main() {

	config.Get("config.json")
	assets.Get()

	ebiten.SetWindowTitle(config.General.WindowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)

	if config.General.Fullscreen {
		ebiten.SetFullscreen(true)
		config.General.WindowSizeX, config.General.WindowSizeY = ebiten.ScreenSizeInFullscreen()
	} else {
		ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)
	}

	if config.General.SpawnOnMouse {
		ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	}

	g := game{system: particles.NewSystem()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
