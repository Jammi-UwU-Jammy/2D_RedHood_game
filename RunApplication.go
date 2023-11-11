package main

import (
	"RedHood/gameImpl"
	"RedHood/util"
	"embed"
)

//go:embed assets/*
var EmbededAssets embed.FS

func main() {
	util.Assets = EmbededAssets
	gameApp := gameImpl.NewGame()
	gameImpl.RunGame(gameApp)

}
