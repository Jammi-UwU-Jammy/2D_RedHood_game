package environments

import (
	"RedHood/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	_ "github.com/lafriks/go-tiled"
)

const (
	tiledBG       = "assets/bground/default.tmx"
	tiledObstacle = "bground/default.tmx"
)

type BGround struct {
	image *ebiten.Image
	ops   *ebiten.DrawImageOptions

	tiledHash    map[uint32]*ebiten.Image
	obstacleHash map[uint32]*ebiten.Image
	Level        *tiled.Map
}

func NewDefaultBackground() *BGround {
	file := tiled.WithFileSystem(util.Assets)
	bgMap, err := tiled.LoadFile(tiledBG, file)
	util.CheckErrExit(100, err)
	//obstacleMap, err := tiled.LoadFile(tiledObstacle, file)
	//util.CheckErrExit(101, err)

	ops := &ebiten.DrawImageOptions{}
	bg := BGround{ops: ops}

	bg.Level = bgMap
	bg.tiledHash = util.MakeEImagesFromMap(*bgMap)
	//bg.obstacleHash = util.MakeEImagesFromMap(*obstacleMap)

	return &bg
}

func (bg *BGround) Update() {

}

func (bg *BGround) Draw(screen *ebiten.Image) {
	bg.rendALayer(0, screen)
	bg.rendALayer(1, screen)
	//bg.rendALayer(2, screen)
}

func (bg *BGround) rendALayer(layer int, screen *ebiten.Image) {
	for tileY := 0; tileY < bg.Level.Height; tileY += 1 {
		for tileX := 0; tileX < bg.Level.Width; tileX += 1 {
			bg.ops.GeoM.Reset()

			TileXpos := float64(bg.Level.TileWidth * tileX)
			TileYpos := float64(bg.Level.TileHeight * tileY)
			bg.ops.GeoM.Translate(TileXpos, TileYpos)

			tileToDraw := bg.Level.Layers[layer].Tiles[tileY*bg.Level.Width+tileX]
			if tileToDraw.ID != 0 {
				bgTileToDraw := bg.tiledHash[tileToDraw.ID]
				screen.DrawImage(bgTileToDraw, bg.ops)
			}
		}
	}
}
