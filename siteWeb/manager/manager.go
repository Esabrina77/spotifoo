package manager

//struct pour stocker les infos de l'album

type Album struct {
	Artist           Artists      `json:"artists"`
	Name             string       `json:"name"`
	Images           []Images     `json:"images"`
	ReleaseDate      string       `json:"release_date"`
	TotalTracks      int          `json:"total_tracks"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	AvailableMarkets []string     `json:"available_markets"`
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
type AlbumResponse struct {
	Albums []Album `json:"albums"`
}

//struct pour stocker les infos d'une piste

type Tracks struct {
	Name         string       `json:"name"`
	AlbumName    string       `json:"album_name"`
	Artist       Artists      `json:"artists"`
	ReleaseDate  string       `json:"release_date"`
	ExternalUrls ExternalUrls `json:"external_urls"`
}

// func PrintColorResult(color string, message string) {
// 	colorCode := ""
// 	switch color {
// 	case "red":
// 		colorCode = "\033[31m"
// 	case "green":
// 		colorCode = "\033[32m"
// 	case "yellow":
// 		colorCode = "\033[33m"
// 	case "blue":
// 		colorCode = "\033[34m"
// 	case "purple":
// 		colorCode = "\033[35m"

// 	default: //REMETTRE LA COULEUR INITIALE (blanc)
// 		colorCode = "\033[0m"
// 	}
// 	fmt.Printf("%s%s\033[0m", colorCode, message)
// }
