package goflappy

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	imagePipe *ebiten.Image
)

type Pipe struct {
	position      Position
	pipeImage     *ebiten.Image
	scored        bool
	width, height float64
}

func init() {
	var err error
	imagePipe, _, err = ebitenutil.NewImageFromFile("assets/pipe.png")
	if err != nil {
		log.Fatal(err)
	}
}

func NewPipe(position Position) Pipe {
	return Pipe{
		position:  position,
		pipeImage: imagePipe,
		width:     50,
		height:    300,
	}
}

func (p *Pipe) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(p.position.X, p.position.Y+pipeGap)
	screen.DrawImage(p.pipeImage, op)

	op.GeoM.Reset()
	op.GeoM.Rotate(math.Pi)
	op.GeoM.Translate(p.position.X+50, p.position.Y-pipeGap)
	screen.DrawImage(p.pipeImage, op)
}

func (p *Pipe) TopPipe() Rect {
	return Rect{
		position: Position{
			X: p.position.X,
			Y: p.position.Y - pipeGap - p.height,
		},
		Width:  p.width,
		Height: p.height,
	}
}

func (p *Pipe) BottomPipe() Rect {
	return Rect{
		position: Position{
			X: p.position.X,
			Y: p.position.Y + pipeGap,
		},
		Width:  p.width,
		Height: p.height,
	}
}
