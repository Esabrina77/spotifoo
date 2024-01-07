package manager

import "fmt"

//struct pour stocker les infos de l'album

type AlbumsData struct {
	Href  string `json:"href"`
	Items []struct {
		AlbumGroup string `json:"album_group"`
		AlbumType  string `json:"album_type"`
		Artists    []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			Id   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			Uri  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		ExternalUrls     struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string `json:"href"`
		Id     string `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			Url    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name                 string `json:"name"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		TotalTracks          int    `json:"total_tracks"`
		Type                 string `json:"type"`
		Uri                  string `json:"uri"`
	} `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int         `json:"total"`
}

type Artists struct {
	Name string `json:"name"`
}
type ExternalUrls struct {
	Spotify string `json:"spotify"`
}
type Images struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

// type AlbumResponse struct {
// 	Albums []Album `json:"albums"`
// }

//struct pour stocker les infos d'une piste

type Tracks struct {
	Name         string       `json:"name"`
	AlbumName    string       `json:"album_name"`
	Artist       Artists      `json:"artists"`
	ReleaseDate  string       `json:"release_date"`
	ExternalUrls ExternalUrls `json:"external_urls"`
}

func PrintColorResult(color string, message string) {
	colorCode := ""
	switch color {
	case "red":
		colorCode = "\033[31m"
	case "green":
		colorCode = "\033[32m"
	case "yellow":
		colorCode = "\033[33m"
	case "blue":
		colorCode = "\033[34m"
	case "purple":
		colorCode = "\033[35m"

	default: //REMETTRE LA COULEUR INITIALE (blanc)
		colorCode = "\033[0m"
	}
	fmt.Printf("%s%s\033[0m", colorCode, message)
}
