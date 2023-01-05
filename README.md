# R1.01.SAE.eq01_AMBROISE-Thomas_GOUIN-Charly



## Pré-requis
> Golang (à installer sur le site https://go.dev)

## Installation
Installer le code source sur gitlab directement ou en passant par par le terminal avec la commande suivante :<br>
```
git clone https://gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae/projets/groupe2/r1.01.sae.eq01_ambroise-thomas_gouin-charly.git
```
Une fois installer, il faudra ce rendre dans le terminal et aller dans le dossier installer auparavant et executer les commandes suivantes :
```
go install
go build
```
Et ensuite vous pourrez lancer le fichier 'project-particles.exe' pour lancer la génération de particule.

## Config.json
> - WindowTitle | Choisir le nom de la fenêtre de l'application.
> - WindowSizeX | Longueur de la fenêtre (en pixel).
> - WindowSizeY | Largeur de la fenêtre (en pixel).
> - ParticleImage | L'image afficher pour les particules (chemin)
> - Debug | Activer l'affichage des informations en haut à gauche de l'écran (true/false)
> - SpawnRate | Nombre de particule par apparition
> - SpawnFrameRate | Fréquence d'apparition des nouvelles particules (60 = 1 second)
> - Fullscreen | Choisir si l'application se lance en plein écran
> - ParticleMovementPreset | Mouvement de particule préfait, les presets : "falling"
> - InitNumParticles | Le nombre de particule à l'écran au démarrage
> - RandomSpawn | Apparition des particules à un endroit aléatoire (true/false)
> - RandomColor | Apparition des particules avec une couleur aléatoire (true/false)
> #### Seulement quand RandomColor est false 
> - ColorR | Gestion de la nuance de Rouge des particules (de 0.0 à 1.0)
> - ColorG | Gestion de la nuance de Vert des particules (de 0.0 à 1.0)
> - ColorB | Gestion de la nuance de Bleu des particules (de 0.0 à 1.0)
> #### Seulement quand RandomOpacity est false
> - Opacity | Gestion de la Transparence des particules (de 0.0 à 1.0)
> #### Seulement quadn RandomSpawn est false
> - SpawnX | Gestion de la coordonnée X des particules au spawn
> - SpawnY | Gestion de la coordonnée Y des particules au spawn