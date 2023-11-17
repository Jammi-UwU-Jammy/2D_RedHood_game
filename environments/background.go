package environments

import (
	"RedHood/characters"
	"RedHood/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	_ "github.com/lafriks/go-tiled"
)

const (
	STD_TILE_WIDTH = 32
	tiledBG        = "assets/bground/default.tmx"
	tiledObstacle  = "bground/default.tmx"
)

type Map struct {
	image *ebiten.Image
	ops   *ebiten.DrawImageOptions

	tiledHash    map[uint32]*ebiten.Image
	obstacleHash map[uint32]*ebiten.Image
	TiledMap     *tiled.Map

	Enemies []*characters.Mob
}

func NewDefaultBackground(enemies []*characters.Mob) *Map {
	file := tiled.WithFileSystem(util.Assets)
	bgMap, err := tiled.LoadFile(tiledBG, file)
	util.CheckErrExit(100, err)

	ops := &ebiten.DrawImageOptions{}
	bg := Map{ops: ops}

	bg.TiledMap = bgMap
	bg.tiledHash = util.MakeEImagesFromMap(*bgMap)
	bg.Enemies = enemies

	return &bg
}

func (bg *Map) Update() {
	//for _, mob := range bg.Enemies {
	//	mob.Update()
	//}
}

func (bg *Map) PopulateMobs(mobs []*characters.Mob) {
	bg.Enemies = mobs
}

func (bg *Map) Draw(screen *ebiten.Image) {
	bg.rendALayer(0, screen)
	bg.rendALayer(1, screen)

	bg.renderMobs(screen)
	//bg.rendALayer(2, screen)
}

func (bg *Map) rendALayer(layer int, screen *ebiten.Image) {
	for tileY := 0; tileY < bg.TiledMap.Height; tileY += 1 {
		for tileX := 0; tileX < bg.TiledMap.Width; tileX += 1 {
			bg.ops.GeoM.Reset()

			TileXpos := float64(bg.TiledMap.TileWidth * tileX)
			TileYpos := float64(bg.TiledMap.TileHeight * tileY)

			tileToDraw := bg.TiledMap.Layers[layer].Tiles[tileY*bg.TiledMap.Width+tileX]
			if tileToDraw.ID != 0 {
				bgTileToDraw := bg.tiledHash[tileToDraw.ID]
				var tileAdjustY float64 = 0
				if layer == 1 {
					tileAdjustY = float64(bgTileToDraw.Bounds().Dy()) - STD_TILE_WIDTH
				}
				bg.ops.GeoM.Translate(TileXpos, TileYpos-tileAdjustY) //-tileAdjustY

				screen.DrawImage(bgTileToDraw, bg.ops)
			}
		}
	}
}

func (bg *Map) renderMobs(screen *ebiten.Image) {
	for _, mob := range bg.Enemies {
		mob.Draw(screen)
	}
}
