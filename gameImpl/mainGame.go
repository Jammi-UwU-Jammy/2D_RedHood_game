package gameImpl

import (
	"RedHood/characters"
	"RedHood/environments"
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

type Game struct {
	bagUI    *ebitenui.UI
	playerUI *ebitenui.UI

	screen     *ebiten.Image
	player     *characters.Player
	enemies    []*characters.Mob
	background *environments.Map

	obstacles []*tiled.Object
}

func NewGame(player *characters.Player, gameMap *environments.Map) *Game {
	ebiten.SetWindowSize(1600, 960)
	ebiten.SetWindowTitle("Red In Da Hood")

	game := Game{}

	game.player = player
	game.bagUI = environments.NewUI()
	game.playerUI = environments.PlayerUI()

	game.background = gameMap
	game.obstacles = game.background.TiledMap.ObjectGroups[0].Objects
	game.enemies = game.background.Enemies

	return &game
}

func (g Game) Update() error {

	g.background.Update()
	g.player.Update(g.obstacles)

	g.bagUI.Update()
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%4.2f:%4.2f", g.player.LocX, g.player.LocY))
	g.player.Draw(screen)

	g.playerUI.Draw(screen)
	g.bagUI.Draw(screen)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
