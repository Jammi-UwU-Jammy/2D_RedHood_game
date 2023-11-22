package util

import (
	"embed"
	"fmt"
	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"golang.org/x/mobile/asset"
	"log"
)

var Assets embed.FS
var SfxContext = audio.NewContext(48000)

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

func LoadEmbededSound(filePath string, errCode int) *audio.Player {
	embededFile, err := asset.Open(filePath)
	CheckErrExit(errCode, err)

	stream, err := wav.Decode(SfxContext, embededFile)
	CheckErrExit(11, err)
	sound, err := audio.NewPlayer(SfxContext, stream)
	CheckErrExit(12, err)
	return sound
}

func PlaySound(s *audio.Player) {
	if !s.IsPlaying() {
		s.Rewind()
		s.Play()
	}
}

func MakeEImagesFromMap(tiledMap tiled.Map) map[uint32]*ebiten.Image {
	idToImage := make(map[uint32]*ebiten.Image)
	for _, tile := range tiledMap.Tilesets[0].Tiles {
		mSource := "assets/bground/" + tile.Image.Source
		ebitenImageTile, _, err := ebitenutil.NewImageFromFile(mSource)
		CheckErrExit(10, err)
		idToImage[tile.ID] = ebitenImageTile
	}
	return idToImage
}

func ImageNineSlice(img *ebiten.Image, centerWidth int, centerHeight int) *image.NineSlice {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	return image.NewNineSlice(img,
		[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
		[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight})
}
