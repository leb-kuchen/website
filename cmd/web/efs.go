package web

import "embed"

//go:embed "js"
var JSFiles embed.FS

//go:embed "css"
var CSSFiles embed.FS
