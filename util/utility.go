package util

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"golang.org/x/mobile/asset"
	"log"
)

var Assets embed.FS
var SfxContext = audio.NewContext(48000)

type Point struct {
	X int
	Y int
}

func CheckErrExit(errCode int, err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatalf("error occurred, code: %d\n", errCode)
	}
}

func LoadEmbeddedImage(filePath string, errCode int) *ebiten.Image {
	embeddedFile, err := asset.Open(filePath)
	CheckErrExit(errCode, err)
	ebitenImage, _, err := ebitenutil.NewImageFromReader(embeddedFile)
	CheckErrExit(errCode, err)
	return ebitenImage
}

func LoadEmbededSound(filePath string, errCode int) asset.File {
	embededFile, err := asset.Open(filePath)
	CheckErrExit(errCode, err)
	return embededFile
}

func MakeEImagesFromMap(tiledMap tiled.Map) map[uint32]*ebiten.Image {
	idToImage := make(map[uint32]*ebiten.Image)
	for _, tile := range tiledMap.Tilesets[0].Tiles {
		ebitenImageTile, _, err := ebitenutil.NewImageFromFile(tile.Image.Source)
		CheckErrExit(10, err)
		idToImage[tile.ID] = ebitenImageTile
	}
	return idToImage
}
