package data

import "embed"

//go:embed locales/*.json
var Locales embed.FS
