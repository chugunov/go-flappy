package goflappy

import (
	"log"
	"math/rand/v2"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	gameTitle         = "Go Flappy!"
	gravity           = 0.1
	playerStartX      = 100
	pipeSpeed         = 2
	pipeSpawnInterval = 100
	pipeGap           = 60.0
	debugMsgOffsetX   = 10
	debugMsgOffsetY   = 10
)

var (
	imageBackground *ebiten.Image
)

func init() {
	var err error
	imageBackground, _, err = ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	screenHeight float64
	screenWidth  float64
	gameOver     bool
	tick         int
	nextPipeTick int
	background   *ebiten.Image
	score        int
	pipes        []Pipe
	player       *Player
}

func NewGame(screenHeight, screenWidth int) (*Game, string) {
	return &Game{
		gameOver:     false,
		screenHeight: float64(screenHeight),
		screenWidth:  float64(screenWidth),
		tick:         0,
		nextPipeTick: 0,
		background:   imageBackground,
		score:        0,
		pipes:        []Pipe{},
		player:       NewPlayer(Position{X: 100, Y: float64(screenHeight) / 4}),
	}, gameTitle
}

func (g *Game) updatePipes() {
	if g.tick >= g.nextPipeTick {
		y := g.screenWidth/3 + rand.Float64()*g.screenWidth/3
		g.pipes = append(g.pipes, NewPipe(Position{X: g.screenHeight, Y: y}))
		g.nextPipeTick = g.tick + 100 + rand.IntN(100)
	}

	for i := 0; i < len(g.pipes); i++ {
		g.pipes[i].position.X -= pipeSpeed

		if g.pipes[i].position.X < -100 {
			g.pipes = append(g.pipes[:i], g.pipes[i+1:]...)
			i--
		}
	}
}

func (g *Game) GameOver(sound bool) {
	g.gameOver = true
	if sound {
		g.player.SoundHit()
	}
}

func (g *Game) checkCollisions() {
	if g.player.CollidesWithScreen(g.screenHeight) {
		g.GameOver(false)
		return
	}

	for i := range g.pipes {
		if g.player.CollidesWith(g.pipes[i].BottomPipe()) || g.player.CollidesWith(g.pipes[i].TopPipe()) {
			g.GameOver(true)
		}

		if g.player.position.X > g.pipes[i].position.X+g.pipes[i].width && !g.pipes[i].scored {
			g.score++
			g.pipes[i].scored = true
		}
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyR)) {
			g.Restart()
		}
		return nil
	}
	g.tick++
	g.player.Update()
	g.checkCollisions()
	g.updatePipes()

	return nil
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	msg := "TFPS: " + strconv.FormatFloat(ebiten.ActualFPS(), 'f', 2, 64) + "\nScores: " + strconv.Itoa(g.score)
	ebitenutil.DebugPrintAt(screen, msg, debugMsgOffsetX, debugMsgOffsetY)
}

func (g *Game) draweGameOver(screen *ebiten.Image) {
	if !g.gameOver {
		return
	}
	msg := "Game Over! Press R to restart."
	ebitenutil.DebugPrintAt(screen, msg, int(g.screenHeight)/2-100, int(g.screenWidth)/2)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
	g.player.Draw(screen)
	for _, pipe := range g.pipes {
		pipe.Draw(screen)
	}
	g.drawDebug(screen)
	g.draweGameOver(screen)
}

func (g *Game) Layout(outsideHeight, outsideWidth int) (int, int) {
	return int(g.screenHeight), int(g.screenWidth)
}

func (g *Game) Restart() {
	g.gameOver = false
	g.tick = 0
	g.nextPipeTick = 0
	g.score = 0
	g.player.velocity = 0
	g.player.position.X = 100
	g.player.position.Y = g.screenHeight / 4
	g.pipes = []Pipe{}
}
