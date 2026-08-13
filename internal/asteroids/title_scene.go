package asteroids

import (
	"asteroids/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TitleScene struct{}

func NewTitleScene() *TitleScene {
	return &TitleScene{}
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	textToDraw := "1 coin 1 play"
	opts := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}
	opts.ColorScale.ScaleWithColor(color.White)
	opts.GeoM.Translate(float64(screenWidth/2), float64(screenHeight-200))

	text.Draw(
		screen,
		textToDraw,
		&text.GoTextFace{Source: assets.TitleFont, Size: 48},
		opts,
	)
}

func (s *TitleScene) Update(state *State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoToScene(NewGameScene())
	}

	return nil
}
