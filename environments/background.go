package environments

import (
	"RedHood/characters"
	"RedHood/util"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	_ "github.com/lafriks/go-tiled"
	path "github.com/solarlune/paths"
	"strings"
)

const (
	STD_TILE_WIDTH  = 32
	defaultTiledMap = "assets/bground/default.tmx"
	lakeTiledMap    = "assets/bground/forest.tmx"
)

type Map struct {
	image *ebiten.Image
	ops   *ebiten.DrawImageOptions

	tiledHash    map[uint32]*ebiten.Image
	obstacleHash map[uint32]*ebiten.Image
	TiledMap     *tiled.Map
	PathGrid     *path.Grid

	Enemies []*characters.Mob
}

func NewBackground(tmxPath string) *Map {
	file := tiled.WithFileSystem(util.Assets)
	bgMap, err := tiled.LoadFile(tmxPath, file)
	util.CheckErrExit(100, err)

	ops := &ebiten.DrawImageOptions{}
	bg := Map{ops: ops}

	bg.TiledMap = bgMap
	bg.tiledHash = util.MakeEImagesFromMap(*bgMap)

	pathString := makeSearchMap(bg.TiledMap)
	fmt.Println(len(pathString))
	searchPthMp := path.NewGridFromStringArrays(pathString, bg.TiledMap.TileWidth, bg.TiledMap.TileHeight)
	searchPthMp.SetWalkable('1', true)
	searchPthMp.SetWalkable('0', false)

	bg.PathGrid = searchPthMp

	return &bg
}

func NewDefaultMap() *Map {
	bg := NewBackground(defaultTiledMap)
	return bg
}

func NewLakeMap() *Map {
	bg := NewBackground(lakeTiledMap)
	return bg
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
	bg.rendALayer(1, screen)
	//bg.rendALayer(2, screen)

	bg.renderMobs(screen)
	bg.rendALayer(2, screen)
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
				if layer == 2 {
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
		//fmt.Println()
	}
}

func (bg *Map) SetMobs(mobs []*characters.Mob) {
	bg.Enemies = mobs
}

func makeSearchMap(tiledMap *tiled.Map) []string {
	mapAsStringSlice := make([]string, 0, tiledMap.Height) //each row will be its own string
	row := strings.Builder{}

	for position, tile := range tiledMap.Layers[0].Tiles {
		if position%tiledMap.Width == 0 && position > 0 { // we get the 2d array as an unrolled one-d array
			mapAsStringSlice = append(mapAsStringSlice, row.String())
			row = strings.Builder{}
		}
		if tile.ID == 0 {
			row.WriteString(fmt.Sprintf("%d", 1))
		} else {
			row.WriteString(fmt.Sprintf("%d", 0))
		}
	}
	mapAsStringSlice = append(mapAsStringSlice, row.String())
	for _, i := range mapAsStringSlice {
		fmt.Println(i)
	}
	return mapAsStringSlice
}
