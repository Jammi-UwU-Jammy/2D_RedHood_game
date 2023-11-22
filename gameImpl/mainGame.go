package gameImpl

import (
	"RedHood/characters"
	"RedHood/environments"
	"RedHood/etc"
	"RedHood/util"
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"time"
)

const (
	SFX_PORTAL     = "player/sfx/Retro Blop StereoUP 04.wav"
	SFX_HEAL       = "player/sfx/Retro Charge Magic 11.wav"
	SFX_EQUIP      = "player/sfx/Retro Beeep 20.wav"
	SFX_QUESTCOMPL = "player/sfx/Retro PowerUP StereoUP 05.wav"
)

type Game struct {
	QuestUI  *ebitenui.UI
	playerUI *environments.PlayerUI

	screen     *ebiten.Image
	player     *characters.Player
	enemies    []*characters.Mob
	background *environments.Map

	portals        []*tiled.Object
	universalItems []*etc.Item
	quests         []*etc.Quest

	portalSound *audio.Player
	healSound   *audio.Player
	equipSound  *audio.Player
	questSound  *audio.Player

	//Field for Manager to access
	PlayerDataToSend map[string]interface{}
}

func NewGame(player *characters.Player, gameMap *environments.Map) *Game {
	ebiten.SetWindowSize(1600, 960)
	ebiten.SetWindowTitle("Red In Da Hood")

	game := Game{}

	game.player = player
	//game.QuestUI = environments.NewGridContainer(9)
	game.playerUI = environments.NewPlayerUI()
	game.QuestUI = environments.NewQuestUI()

	game.loadEnvironmentSound()
	game.PopulateQuests()
	game.background = gameMap
	game.portals = game.background.TiledMap.ObjectGroups[0].Objects
	game.enemies = game.background.Enemies
	game.PlayerDataToSend = make(map[string]interface{})

	return &game
}

func (g *Game) Update() error {

	g.background.Update()
	g.PlayerDataToSend = g.player.Update(g.background.TiledMap, g.portals)

	g.playerUI.HP.Configure(widget.ProgressBarOpts.Values(0, g.player.MaxStat.HP, g.player.HP))
	g.playerUI.Update()
	g.UpdateBag()
	g.UpdateQuests()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("HP: %d/%d | ATK: %d",
		g.player.HP, g.player.MaxStat.HP, g.player.MaxStat.ATK))
	g.player.Draw(screen)

	g.DrawEtcItems(screen)
	g.playerUI.Draw(screen)
	g.QuestUI.Draw(screen)
}

func (g *Game) UpdateBag() {
	g.playerUI.Bag.RemoveChildren()
	for _, it := range g.player.Bag {
		listenedItem := it
		container := widget.NewContainer(
			widget.ContainerOpts.BackgroundImage(util.ImageNineSlice(it.GetImage(), 32, 32)),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(32, 32),
			),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MouseButtonPressedHandler(func(args *widget.WidgetMouseButtonPressedEventArgs) {
					//TODO: add attributes
					fmt.Println(listenedItem.Buffs.HP, listenedItem.Buffs.ATK)
					g.player.EquipItem(listenedItem)
					util.PlaySound(g.equipSound)
				}),
			),
		)
		g.playerUI.Bag.AddChild(container)
	}
}

func (g *Game) UpdateQuests() {
	g.QuestUI.Container.RemoveChildren()
	var tempQuest []*etc.Quest
	for _, it := range g.quests {
		q := it
		if q.Conditions() != true {
			environments.CreateAQuest(q.Title, q.Description, g.QuestUI.Container)
			tempQuest = append(tempQuest, q)
		}
	}
	g.quests = tempQuest
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

func (g *Game) loadEnvironmentSound() {
	g.healSound = util.LoadEmbededSound(SFX_HEAL, 320)
	g.equipSound = util.LoadEmbededSound(SFX_EQUIP, 321)
	g.questSound = util.LoadEmbededSound(SFX_QUESTCOMPL, 322)
	g.portalSound = util.LoadEmbededSound(SFX_PORTAL, 323)
}

func (g *Game) PopulateQuests() {
	q := etc.Quest{
		Title:      "Play for more than 10s.",
		Conditions: nil,
		Created:    time.Now(),
	}
	cond := func() bool {
		q.Description = fmt.Sprintf("Playing: %.f/%ds", time.Since(q.Created).Seconds(), 10)
		if util.IsCDExceeded(10, q.Created) {
			util.PlaySound(g.questSound)
			return true
		}
		return false
	}
	q.Conditions = cond

	q1 := etc.Quest{
		Title:       "Loot 3 item.",
		Description: "",
		Conditions:  nil,
	}
	cond1 := func() bool {
		q1.Description = fmt.Sprintf("Items looted: %d/%d", len(g.player.Bag), 3)
		if len(g.player.Bag) >= 3 {
			util.PlaySound(g.questSound)
			return true
		}
		return false
	}
	q1.Conditions = cond1
	g.quests = append(g.quests, &q, &q1)
}
