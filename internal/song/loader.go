package song

import (
	"encoding/json"
	"os"
)

var Songs []Song

func LoadSongs(path string) error {

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &Songs)
}