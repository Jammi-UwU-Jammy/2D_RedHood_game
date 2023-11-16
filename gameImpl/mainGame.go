package gameImpl

import (
	"RedHood/characters"
	"RedHood/environments"
	"RedHood/util"
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

func NewGame() *Game {
	ebiten.SetWindowSize(1600, 960)
	ebiten.SetWindowTitle("Red In Da Hood")

	game := Game{}
	game.bagUI = environments.NewUI()
	game.playerUI = environments.PlayerUI()
	game.background = environments.NewDefaultBackground()
	game.player = characters.NewPlayer()
	game.enemy = characters.NewEnemyMage()

	game.obstacles = game.background.TiledMap.ObjectGroups[0].Objects

	return &game
}

func RunGame(game *Game) {
	err := ebiten.RunGame(game)
	util.CheckErrExit(-1, err)
}

type Game struct {
	bagUI    *ebitenui.UI
	playerUI *ebitenui.UI

	screen     *ebiten.Image
	player     *characters.Player
	enemy      *characters.Mob
	background *environments.BGround

	obstacles []*tiled.Object
}

func (g Game) Update() error {

	g.background.Update()
	g.player.Update(g.obstacles)
	g.enemy.Update(g.player)

	g.bagUI.Update()
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%4.2f:%4.2f", g.player.LocX, g.player.LocY))
	g.player.Draw(screen)
	g.enemy.Draw(screen)

	g.playerUI.Draw(screen)
	//g.bagUI.Draw(screen)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
