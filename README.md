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
> ### <ins>Configuration de la fenêtre</ins>
> - WindowTitle | Choisir le nom de la fenêtre.
> - Debug | Activer l'affichage des informations en haut à gauche de l'écran (true/false)
> - Fullscreen | Choisir si l'application se lance en plein écran
> #### Seulement quand Fullscreen est false
> - WindowSizeX | Longueur de la fenêtre (en pixel)
> - WindowSizeY | Largeur de la fenêtre (en pixel)
> ### <ins>Configuration des particules</ins>
> - ParticleImage | L'image afficher pour les particules (chemin)
> - InitNumParticles | Le nombre de particule à l'écran au démarrage
> - MaxNumParticles | Le nombre maximum de particule à l'écran
> - SpawnRate | Nombre de particule par apparition
> - TimeToSpawn | Fréquence d'apparition des nouvelles particules (60 = 1 second)
> - EnableLifeSpan | Activer la durée de vie des particules
> - LifeSpan | Durée de vie des particules, seulement si EnableLifeSpan est true (60 = 1 second)
> - SmoothSuppression | Supprime les particules d'une manière plus jolie
> - RandomSpawn | Apparition des particules à un endroit aléatoire (true/false)
> - RandomColor | Apparition des particules avec une couleur aléatoire (true/false)
> - RandomOpacity | Apparition des particules avec une opacité aléatoire (true/false)
> - RandomScale | Apparition des particules avec une taille aléatoire (true/false)
> - RandomSpeed | Apparition des particules avec une vitesse aléatoire (true/false)
> #### Seulement quand RandomSpawn est false
> - SpawnX | Gestion de la coordonnée X des particules au spawn (de 0 à longueur de la fenêtre en pixel)
> - SpawnY | Gestion de la coordonnée Y des particules au spawn (de 0 à largeur de la fenêtre en pixel)
> #### Seulement quand RandomColor est false 
> - ColorR | Gestion de la nuance de Rouge des particules (de 0.0 à 1.0)
> - ColorG | Gestion de la nuance de Vert des particules (de 0.0 à 1.0)
> - ColorB | Gestion de la nuance de Bleu des particules (de 0.0 à 1.0)
> #### Seulement quand RandomOpacity est false
> - Opacity | Gestion de la Transparence des particules (de 0.0 à 1.0)
> #### Seulement quand RandomScale est false
> - ScaleX | Gestion de la longueur des particules (de 0.0 à 1.0)
> - ScaleY | Gestion de la largeur des particules (de 0.0 à 1.0)
> #### Seulement quand RandomSpeed est false
> - SpeedX | Gestion de la vitesse horizontale des particules (minimum 1)
> - SpeedY | Gestion de la vitesse verticale des particules (minimum 1)
> #### Seulement quand RandomSpeed est True
> - MaxSpeedX | Gestion de la vitesse horizontale des particules (minimum 1)
> - MaxSpeedY | Gestion de la vitesse verticale des particules (minimum 1)