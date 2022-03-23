package teora

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var msgs = []string{
	"This is a demo build.",
	"project teora is neither affiliated with nor endorsed by GeoEXE.",
}

func scale(res int) int {
	return int(float64(res) * ebiten.DeviceScaleFactor())
}

// Game holds the main state of the game.
type Game struct {
	scroll *Scroll
}

// NewGame creates a game and initalises it's state.
func NewGame() *Game {
	return &Game{
		scroll: nil,
	}
}

// Update updates the game's state.
func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// skip the text if it's still scrolling, otherwise go to the next text.
		if !g.scroll.Done() {
			g.scroll.Skip()
		} else {
			g.scroll.Next()
		}
	}

	return nil
}

// Draw draws to the screen based on the game's state.
func (g *Game) Draw(screen *ebiten.Image) {
	// draw tps/fps at the top left of the screen
	hack.Write(
		fmt.Sprintf("tps: %0.2f", ebiten.CurrentTPS()),
		screen,
		image.Pt(0, 0),
		AlignRight|AlignBottom,
	)

	bounds := teoran.WriteCenter("Hello World!", screen)

	// we can only init scroll here because we need to know the screen size
	if g.scroll == nil {
		g.scroll = NewScroll(
			hack,
			msgs,
			image.Pt(center(screen).X, bounds.Max.Y+20),
			AlignCenter,
		)
	}

	g.scroll.Render(screen)
}

// Layout returns the screen's size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scale(outsideWidth), scale(outsideHeight)
}
