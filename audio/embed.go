package audio

import (
	_ "embed"
)

var (
	//go:embed jump.ogg
	Jump_ogg []byte

	//go:embed hit.ogg
	Hit_ogg []byte
)
