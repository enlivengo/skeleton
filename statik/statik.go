package statik

import (
	"github.com/rakyll/statik/fs"
)

func init() {
	data := "PK\x05\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	fs.Register(data)
}
