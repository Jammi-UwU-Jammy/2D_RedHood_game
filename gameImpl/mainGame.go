package gameImpl

import (
	"RedHood/characters"
	"RedHood/environments"
	"RedHood/util"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2"
)

func NewGame() *Game {
	ebiten.SetWindowSize(1500, 800)
	ebiten.SetWindowTitle("Red In Da Hood")

	game := Game{}
	game.background = environments.NewDefaultBackground()
	game.player = characters.NewPlayer()
	game.enemy = characters.NewEnemy()

	return &game
}

func RunGame(game *Game) {
	err := ebiten.RunGame(game)
	util.CheckErrExit(-1, err)
}

type Game struct {
	screen     *ebiten.Image
	player     *characters.Player
	enemy      *characters.Mob
	background *environments.BGround
}

func (g Game) Update() error {
	g.background.Update()
	g.player.Update()
	g.enemy.Update()

	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	g.player.Draw(screen)
	g.enemy.Draw(screen)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
