package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func animePlanetData() (AnimePlanetData, error) {
	client := &http.Client{}

	fetch := func(path string, v interface{}) error {
		req, err := http.NewRequest(http.MethodGet, "https://www.anime-planet.com/api/export"+path+animePlanetClient.userId, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Cookie", animePlanetClient.cookie)

		// Mustve been the wind
		req.Header.Set("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36`)

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
	if err := fetch("/anime/", &anime); err != nil {
		return AnimePlanetData{}, err
	}

	var manga AnimePlanetManga
	if err := fetch("/manga/", &manga); err != nil {
		return AnimePlanetData{}, err
	}

	// Return the data
	return AnimePlanetData{
		User:  anime.User,
		Anime: anime,
		Manga: manga,
	}, nil
}
