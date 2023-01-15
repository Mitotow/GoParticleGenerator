package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	Debug                    bool
	Fullscreen               bool
	WindowSizeX, WindowSizeY int
	GameMode                 string
	ParticleImage            string
	InitNumParticles         int
	MaxNumParticles          int
	SpawnRate                float64
	TimeToSpawn              int
	EnableLifeSpan           bool
	LifeSpan                 int
	KillWhenOutOfScreen      bool
	SmoothSuppression        bool
	RandomSpawn              bool
	RandomRotation           bool
	RandomColor              bool
	RandomOpacity            bool
	RandomScale              bool
	RandomSpeed              bool
	RandomLifeSpan           bool
	SpawnX, SpawnY           int
	Rotation                 float64
	ColorR, ColorG, ColorB   float64
	Opacity                  float64
	ScaleX, ScaleY           float64
	SpeedX, SpeedY           float64
	BetterRandomColor        bool
	MaxRotation, MinRotation float64
	MaxOpacity, MinOpacity   float64
	MaxScaleX, MaxScaleY     float64
	MinScaleX, MinScaleY     float64
	MaxSpeedX, MaxSpeedY     float64
	MinSpeedX, MinSpeedY     float64
	MaxLifeSpan, MinLifeSpan int
}

var General Config
