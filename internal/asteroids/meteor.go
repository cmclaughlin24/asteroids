package asteroids

import (
	"asteroids/assets"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationSpeedMin       = -0.02
	rotationSpeedMax       = 0.02
	numberFromLargetMeteor = 4
)

type Meteor struct {
	game          *GameScene
	position      Vector
	rotation      float64
	movement      Vector
	angle         float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewMeteor(baseVelocity float64, g *GameScene, index int) *Meteor {
	// Target center of the screen.
	target := Vector{X: float64(screenWidth) / 2, Y: float64(screenHeight) / 2}

	// Random angle.
	angle := rand.Float64() * 2 * math.Pi

	// The distance from the center that the meteor should spawn at.
	r := float64(screenWidth)/2.0 + 500

	// Position
	pos := Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	// Movement
	velocity := baseVelocity + rand.Float64()*1.5
	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}
	normalizedDirection := direction.Normalize()
	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	return &Meteor{
		game:          g,
		position:      pos,
		movement:      movement,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin),
		sprite:        sprite,
		angle:         angle,
	}
}
