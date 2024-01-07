package assets

import (
	"embed"
	"io/fs"
)

//go:embed front/out
//go:embed front/out/_next
//go:embed front/out/_next/static/**/**/*.js
//go:embed front/out/_next/static/*/*.js
var next embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(next, "front/out")
}
