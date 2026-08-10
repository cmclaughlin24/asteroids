package asteroids

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Player       *Player
	screenWidth  int
	screenHeight int
}

func NewGame(screenWidth, screenHeight int) *Game {
	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (g *Game) Update() error {
	g.Player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}
