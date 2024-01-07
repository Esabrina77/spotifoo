package manager

//struct pour stocker les infos de l'album

type Album struct {
	Name        string `json:"name"`
	CoverImage  string `json:"cover_image"`
	ReleaseDate int    `json:"release_date"`
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
