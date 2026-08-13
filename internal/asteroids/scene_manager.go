package asteroids

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const transitionMax = 25

type Scene interface {
	Update(state *State) error
	Draw(screen *ebiten.Image)
}

type State struct {
	SceneManager *SceneManager
	Input        *Input
}

type SceneManager struct {
	current         Scene
	next            Scene
	transitionFrom  *ebiten.Image
	transitionTo    *ebiten.Image
	transitionCount int
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		transitionFrom: ebiten.NewImage(screenWidth, screenHeight),
		transitionTo:   ebiten.NewImage(screenWidth, screenHeight),
	}
}

func (s *SceneManager) Draw(r *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(r)
		return
	}

	s.transitionFrom.Clear()
	s.current.Draw(s.transitionFrom)

	s.transitionTo.Clear()
	s.next.Draw(s.transitionTo)

	r.DrawImage(s.transitionFrom, nil)

	alpha := 1 - float32(s.transitionCount)/float32(transitionMax)
	opts := &ebiten.DrawImageOptions{}
	opts.ColorScale.ScaleAlpha(alpha)

	r.DrawImage(s.transitionTo, opts)
}

func (s *SceneManager) Update(_ *Input) error {
	if s.transitionCount == 0 {
		return s.current.Update(&State{
			SceneManager: s,
		})
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}

	s.current = s.next
	s.next = nil

	return nil
}

func (s *SceneManager) GoToScene(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMax
	}
}
