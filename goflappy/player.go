package goflappy

import (
	"bytes"
	"fmt"
	"log"
	"math"

	raudio "github.com/chugunov/go-flappy/audio"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	imageGopher  *ebiten.Image
	jumpPlayer   *audio.Player
	hitPlayer    *audio.Player
	audioContext *audio.Context
)

func init() {
	var err error
	imageGopher, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	if audioContext == nil {
		audioContext = audio.NewContext(48000)
	}
	jumpD, err := vorbis.DecodeF32(bytes.NewReader(raudio.Jump_ogg))
	if err != nil {
		log.Fatal(err)
	}
	jumpPlayer, err = audioContext.NewPlayerF32(jumpD)
	if err != nil {
		log.Fatal(err)
	}
	jumpPlayer.SetVolume(0.1)

	hitD, err := vorbis.DecodeF32(bytes.NewReader(raudio.Hit_ogg))
	if err != nil {
		log.Fatal(err)
	}
	hitPlayer, err = audioContext.NewPlayerF32(hitD)
	if err != nil {
		log.Fatal(err)
	}
	hitPlayer.SetVolume(0.1)
}

type Player struct {
	position      Position
	velocity      float64
	gopherImage   *ebiten.Image
	width, height float64
	jumpPlayer    *audio.Player
	hitPlayer     *audio.Player
}

func NewPlayer(position Position) *Player {
	return &Player{
		position:    position,
		gopherImage: imageGopher,
		width:       50,
		height:      50,
		jumpPlayer:  jumpPlayer,
		hitPlayer:   hitPlayer,
	}
}

func (p *Player) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.velocity = -2.5
		p.SoundJump()
	} else {
		p.velocity += gravity
		p.velocity = math.Min(p.velocity, 2.5)
		p.position.Y += p.velocity
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.gopherImage, op)
}

func (p *Player) SoundJump() {
	if err := p.jumpPlayer.Rewind(); err != nil {
		return
	}
	p.jumpPlayer.Play()
}

func (p *Player) SoundHit() {
	if err := p.hitPlayer.Rewind(); err != nil {
		return
	}
	p.hitPlayer.Play()
}

func (p *Player) Rect() Rect {
	return Rect{
		position: Position{
			X: p.position.X,
			Y: p.position.Y,
		},
		Width:  p.width,
		Height: p.height,
	}
}

func (p *Player) CollidesWithScreen(screenHeight float64) bool {
	playerRect := p.Rect()
	return playerRect.position.Y < 0 || playerRect.position.Y+playerRect.Height > screenHeight
}

func (p *Player) CollidesWith(rect Rect) bool {
	if p.Rect().Overlaps(rect) {
		fmt.Println("Collision detected:", p.position, rect)
		return true
	}
	return false
}
