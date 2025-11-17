package main

// new struct
type AnimePlanet struct {
	cookie string
	userId string
	token  string
}

type AnimePlanetUser struct {
	Name string `json:"name"`
}

type AnimePlanetData struct {
	User  AnimePlanetUser  `json:"user"`
	Manga AnimePlanetManga `json:"manga"`
	Anime AnimePlanetAnime `json:"anime"`
}

type AnimePlanetExport struct {
	Type    string `json:"type"`
	Date    string `json:"date"`
	Version string `json:"version"`
}

type AnimePlanetManga struct {
	User AnimePlanetUser `json:"user"`

	Export AnimePlanetExport `json:"export"`

	Entries []struct {
		Name      string  `json:"name"`
		Status    string  `json:"status"`
		Started   string  `json:"started"`
		Completed *string `json:"completed"`
		Rating    int     `json:"rating"`
		Ch        int     `json:"ch"`
		Vol       int     `json:"vol"`
	} `json:"entries"`
}

type AnimePlanetAnime struct {
	User AnimePlanetUser `json:"user"`

	Export AnimePlanetExport `json:"export"`

	Entries []struct {
		Name      string  `json:"name"`
		Status    string  `json:"status"`
		Started   *string `json:"started"`
		Completed *string `json:"completed"`
		Rating    float64 `json:"rating"`
		Times     int     `json:"times"`
		Eps       int     `json:"eps"`
	} `json:"entries"`
}
