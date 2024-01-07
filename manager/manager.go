package manager

import "fmt"

//struct pour stocker les infos de l'album

type Album struct {
	Name        string `json:"name"`
	CoverImage  string `json:"cover_image"`
	ReleaseDate string `json:"release_date"`
	NumTracks   int    `json:"num_tracks"`
}

//struct pour stocker les infos d'une piste

type Tracks struct {
	Name        string `json:"name"`
	AlbumName   string `json:"album_name"`
	ArtistName  string `json:"artist_name"`
	ReleaseDate string `json:"release_date"`
	ExternalURL string `json:"external_url"`
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
