package assets

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed ui/static
var assets embed.FS

//go:embed ui/html
var templates embed.FS

func Assets() fs.FS {
	staticFs := fs.FS(assets)
	assetsContent, err := fs.Sub(staticFs, "ui/static")
	if err != nil {
		log.Fatal("Konnte static assets nicht öffnen")
	}
	return assetsContent
}

func Templates() embed.FS {
	return templates
}
