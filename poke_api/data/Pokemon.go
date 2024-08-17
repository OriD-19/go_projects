package data

type Pokemon struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`
	} `json:"moves"`
	Sprite struct {
		Other struct {
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
	} `json:"sprites"`
}

type Move struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokeType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Sprite struct {
	OfficialArtwork struct {
		FrontDefault string `json:"front-default"`
	} `json:"official-artwork"`
}
