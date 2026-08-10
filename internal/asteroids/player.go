package asteroids

import (
	"asteroids/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSecond = math.Pi
	maxAcceleration   = 8.0
)

var currentAcceleration float64

type Player struct {
	game     *Game
	sprite   *ebiten.Image
	rotation float64
	position Vector
	velocity float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		game:   game,
		sprite: sprite,
	}

	width, height := p.dimensions()
	halfW := width / 2
	halfH := height / 2
	pos := Vector{
		X: float64(game.screenWidth)/2 - halfW,
		Y: float64(game.screenHeight)/2 - halfH,
	}
	p.position = pos

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	width, height := p.dimensions()
	halfW := width / 2
	halfH := height / 2
	opts := &ebiten.DrawImageOptions{}

	// Rotation
	opts.GeoM.Translate(-halfW, -halfH)
	opts.GeoM.Rotate(p.rotation)
	opts.GeoM.Translate(halfW, halfH)

	// Position
	opts.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, opts)
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

	p.accelerate()
}

func (p *Player) accelerate() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.keepOnScreen()

		if currentAcceleration < maxAcceleration {
			currentAcceleration = p.velocity + 4
		}

		if currentAcceleration >= 8 {
			currentAcceleration = 8
		}

		p.velocity = currentAcceleration

		dx := math.Sin(p.rotation) * currentAcceleration
		dy := math.Cos(p.rotation) * -currentAcceleration

		p.position.X += dx
		p.position.Y += dy
	}
}

func (p *Player) keepOnScreen() {
	if p.position.X >= float64(p.game.screenWidth) {
		p.position.X = 0
	}

	if p.position.X < 0 {
		p.position.X = float64(p.game.screenWidth)
	}

	if p.position.Y >= float64(p.game.screenHeight) {
		p.position.Y = 0
	}

	if p.position.Y < 0 {
		p.position.Y = float64(p.game.screenHeight)
	}
}

func (p Player) dimensions() (float64, float64) {
	bounds := p.sprite.Bounds()
	return float64(bounds.Dx()), float64(bounds.Dy())

}
