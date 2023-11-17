package main

import (
	"RedHood/gameImpl"
	"RedHood/util"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var EmbededAssets embed.FS

func main() {
	util.Assets = EmbededAssets
	//gameApp := gameImpl.NewGame()
	//gameImpl.RunGame(gameApp)

	gameManager := gameImpl.Manager{}
	gameManager.PopulateResources()

	go gameManager.Start()

	RunGame(gameManager.CurrentGame)
	gameManager.Stop()
}

func RunGame(game *gameImpl.Game) {
	err := ebiten.RunGame(game)
	util.CheckErrExit(-1, err)
}
