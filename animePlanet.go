package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func animePlanetData() (AnimePlanetData, error) {
	client := &http.Client{}

	fetch := func(path string, v interface{}) error {
		req, err := http.NewRequest(http.MethodGet, "http://api.scrape.do/", nil)
		if err != nil {
			return err
		}
		req.Header.Set("Accept", "application/json")

		url := path + animePlanetClient.userId

		values := req.URL.Query()
		values.Add("url", url)
		values.Add("token", animePlanetClient.token)
		values.Add("setCookies", "ap="+animePlanetClient.cookie)
		req.URL.RawQuery = values.Encode()

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("anime planet API returned status code %d: %s", resp.StatusCode, resp.Status)
		}

		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(v); err != nil {
			return fmt.Errorf("error decoding response for %s: %w", path, err)
		}
		return nil
	}

	var anime AnimePlanetAnime
	if err := fetch("https://www.anime-planet.com/api/export/anime/", &anime); err != nil {
		return AnimePlanetData{}, err
	}

	var manga AnimePlanetManga
	if err := fetch("https://www.anime-planet.com/api/export/manga/", &manga); err != nil {
		return AnimePlanetData{}, err
	}

	// Return the data
	return AnimePlanetData{
		User:  anime.User,
		Anime: anime,
		Manga: manga,
	}, nil
}
