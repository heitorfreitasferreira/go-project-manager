package web

import "embed"

//go:embed "assets"
var Files embed.FS

var FavIcon = "favicon.ico"
