package convoy

import (
	"embed"
)

//go:embed VERSION
var f embed.FS

func ReadVersion() ([]byte, error) {
	data, err := f.ReadFile("VERSION")
	if err != nil {
		return nil, err
	}

	return data, nil
}
