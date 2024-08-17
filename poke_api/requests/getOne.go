package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"OriD19.com/api/poke/data"
)

func GetOne(id int) (*data.Pokemon, error) {
	req, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id))

	if err != nil {
		return nil, err
	}

	d, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	var p data.Pokemon
	err = json.Unmarshal(d, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
