package asteroids

import "github.com/hajimehoshi/ebiten/v2"

var (
	screenWidth  int
	screenHeight int
)

type Game struct {
	sceneManager *SceneManager
	input        *Input
}

func NewGame(width, height int) *Game {
	screenWidth = width
	screenHeight = height
	return &Game{}
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = NewSceneManager()
		g.sceneManager.GoToScene(NewTitleScene())
	}

	g.input.Update()
	if err := g.sceneManager.Update(g.input); err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}
