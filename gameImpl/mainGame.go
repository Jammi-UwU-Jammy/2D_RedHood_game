package gameImpl

import (
	"RedHood/characters"
	"RedHood/environments"
	"RedHood/etc"
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

type Game struct {
	bagUI    *ebitenui.UI
	playerUI *environments.PlayerUI

	screen     *ebiten.Image
	player     *characters.Player
	enemies    []*characters.Mob
	background *environments.Map

	obstacles      []*tiled.Object
	universalItems []*etc.Item

	//Field for Manager to access
	PlayerData map[string]interface{}
}

func NewGame(player *characters.Player, gameMap *environments.Map) *Game {
	ebiten.SetWindowSize(1600, 960)
	ebiten.SetWindowTitle("Red In Da Hood")

	game := Game{}

	game.player = player
	//game.bagUI = environments.NewUI()
	game.playerUI = environments.NewPlayerUI()

	game.background = gameMap
	game.obstacles = game.background.TiledMap.ObjectGroups[0].Objects
	game.enemies = game.background.Enemies
	game.PlayerData = make(map[string]interface{})

	return &game
}

func (g *Game) Update() error {

	g.background.Update()
	g.PlayerData = g.player.Update(g.obstacles)

	g.playerUI.HP.Configure(widget.ProgressBarOpts.Values(0, 100, g.player.HP))
	g.playerUI.Update()
	//g.bagUI.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d", g.player.HP))
	g.player.Draw(screen)

	g.DrawEtcItems(screen)
	g.playerUI.Draw(screen)
	//g.bagUI.Draw(screen)
}

func (g *Game) DrawEtcItems(screen *ebiten.Image) {
	ops := ebiten.DrawImageOptions{}
	for _, item := range g.universalItems {
		ops.GeoM.Reset()
		ops.GeoM.Translate(item.LocX, item.LocY)
		screen.DrawImage(item.GetImage(), &ops)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
