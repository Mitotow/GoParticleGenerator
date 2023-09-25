# Générateur de particule - Golang
This project was imported from my university gitlab, here are some informations :
* 25 Commits
* 2 Contributors
* Last update : Jan. 18 2023
* Still updated : No

## Pré-requis
> Golang (à installer sur le site https://go.dev)

## Installation
Installer le code source sur gitlab directement ou en passant par par le terminal avec la commande suivante :<br>
```
git clone https://github.com/Mitotow/GoParticleGenerator
```
Une fois installer, il faudra ce rendre dans le terminal et aller dans le dossier installer auparavant et executer les commandes suivantes :
```
go install
go build
```
Et ensuite vous pourrez lancer le fichier 'project-particles.exe' pour lancer la génération de particule.

## Config.json
> ### <ins>Configuration de la fenêtre</ins>
> - WindowTitle | Choisir le nom de la fenêtre.
> - Debug | Activer l'affichage des informations en haut à gauche de l'écran (true/false)
> - Fullscreen | Choisir si l'application se lance en plein écran
> #### Seulement quand Fullscreen est false
> - WindowSizeX | Longueur de la fenêtre (en pixel)
> - WindowSizeY | Largeur de la fenêtre (en pixel)
> ### <ins>Configuration des particules</ins>
> - ParticleImage | L'image afficher pour les particules (chemin vers un .png)
> - InitNumParticles | Le nombre de particule à l'écran au démarrage
> - MaxNumParticles | Le nombre maximum de particule à l'écran
> - SpawnRate | Nombre de particule par apparition
> - TimeToSpawn | Fréquence d'apparition des nouvelles particules (60 = 1 second)
> - EnableLifeSpan | Activer la durée de vie des particules (true/false)
> - LifeSpan | Durée de vie des particules, seulement si EnableLifeSpan est true (60 = 1 second)
> - KillWhenOutOfScreen | Supprimer les particules lorsqu'elles sortent de l'écran (true/false)
> - SmoothSuppression | Supprime les particules d'une manière plus jolie, à utiliser avec le Life Span d'activer (true/false)
> - EnableGravityEffect | Ajoute un effet de gravité aux particules (true/false)
> - SpawnOnMouse | Une fois activer, lorsque vous lancez le programme, si vous maintenez clique gauche, des particules spawnerons à la position de votre curseur (true/false)
> - EnableCollision | Activer le système de collision entre les particules (marche pas très bien avec beaucoup de particule)
> - SpawnCenterX | Le spawn des particules ce fait au centre sur l'axe X (Seulement si Random Spawn = false)
> - SpawnCenterY | Le spawn des particules ce fait au centre sur l'axe Y (Seulement si Random Spawn = false)
> - RandomSpawn | Apparition des particules à un endroit aléatoire (true/false)
> - RandomRotation | Apparition des particules avec une rotation aléatoire (true/false)
> - RandomColor | Apparition des particules avec une couleur aléatoire (true/false)
> - RandomOpacity | Apparition des particules avec une opacité aléatoire (true/false)
> - RandomScale | Apparition des particules avec une taille aléatoire (true/false)
> - RandomSpeed | Apparition des particules avec une vitesse aléatoire (true/false)
> - RandomLifeSpan | Apparition des particules avec une durée de vie aléatoire (true/false)
> #### Seulement quand RandomSpawn est false
> - SpawnX | Gestion de la coordonnée X des particules au spawn (de 0 à longueur de la fenêtre en pixel, -1 = centrer)
> - SpawnY | Gestion de la coordonnée Y des particules au spawn (de 0 à largeur de la fenêtre en pixel, -1 = centrer)
> #### Seulement quand RandomRotation est false
> - Rotation | Degrés de rotation des particules (de 0.0 à 1.0)
> #### Seulement quand RandomColor est false
> - ColorR | Gestion de la nuance de Rouge des particules (de 0.0 à 1.0)
> - ColorG | Gestion de la nuance de Vert des particules (de 0.0 à 1.0)
> - ColorB | Gestion de la nuance de Bleu des particules (de 0.0 à 1.0)
> #### Seulement quand RandomOpacity est false
> - Opacity | Gestion de la Transparence des particules (de 0.0 à 1.0)
> #### Seulement quand RandomScale est false
> - ScaleX | Gestion de la longueur des particules (de 0.0 à 1.0)
> - ScaleY | Gestion de la hauteur des particules (de 0.0 à 1.0)
> #### Seulement quand RandomSpeed est false
> - SpeedX | Gestion de la vitesse horizontale des particules
> - SpeedY | Gestion de la vitesse verticale des particules
> #### Seulement quand RandomColor est true
> - BetterRandomColor | Génère des couleurs mieux visibles (true/false)
> #### Seulement quand RandomRotation est true
> - MaxRotation | Degrés de rotation maximum des particules (de 0.0 à 1.0)
> - MinRotation | Degrés de rotation minimum des particules (de 0.0 à 1.0)
> #### Seulement quand RandomOpacity est true
> - MaxOpacity | Gestion de la transparence maximale des particules (de 0.0 à 1.0)
> - MinOpacity | Gestion de la transparence minimale des particules (de 0.0 à 1.0)
> #### Seulement quand RandomScale est True
> - MaxScaleX | Gestion de la longueur maximale des particules (de 0.0 à 1.0)
> - MinScaleX | Gestion de la longueur horizontale minimale des particules (de 0.0 à 1.0)
> - MaxScaleY | Gestion de la hauteur verticale maximale des particules (de 0.0 à 1.0)
> - MinScaleY | Gestion de la hauteur verticale minimale des particules (de 0.0 à 1.0)
> #### Seulement quand RandomSpeed est True
> - MaxSpeedX | Gestion de la vitesse horizontale maximale des particules
> - MinSpeedX | Gestion de la vitesse horizontale minimale des particules
> - MaxSpeedY | Gestion de la vitesse verticale maximale des particules
> - MinSpeedY | Gestion de la vitesse verticale minimale des particules
> #### Seulement quand RandomLifeSpan est true
> - MaxLifeSpan | Gestion de la durée de vie maximale des particules (60 = 1 second)
> - MinLifeSpan | Gestion de la durée de vie minimale des particules (60 = 1 second)
