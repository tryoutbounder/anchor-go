package sep1

import (
	"net/http"

	"github.com/pelletier/go-toml/v2"
)

func GetToml(url string, httpClient *http.Client) (*StellarToml, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var stellartoml StellarToml
	decoder := toml.NewDecoder(resp.Body)
	err = decoder.Decode(&stellartoml)
	if err != nil {
		return nil, err
	}

	return &stellartoml, nil
}
