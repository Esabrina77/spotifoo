package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"spotifoo/manager"
	initTemplate "spotifoo/templates"
	"time"
)

const Port = "localhost:8080"

// requete Get vers l'api de spotify
// et reception des données obtenues
func MakeApiRequest(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	//gestion du token (en-tete d'authetification)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}

//endoint pour la page des albums de jul

func julAlbumHandler(w http.ResponseWriter, r *http.Request) {

	//requete vers api se spotify pour obtenir
	//les infos des albums de jul

	//ID DE Jul = 3IW7ScrzXmPvZhB27hmfgy
	url := "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy/albums"

	body, err := MakeApiRequest(url)

	if err != nil {
		log.Fatal(err)
		fmt.Println("ERREUR LORS DE LA RECUPERATION DES ALBUMS DE JUL")
	}
	//Analyse des reponses json
	var albums []manager.Album
	err = json.Unmarshal(body, &albums)

	if err != nil {
		log.Fatal(err)
	}

	// Réponse : Charger et exécuter le template avec les données d'albums
	initTemplate.Temp.ExecuteTemplate(w, "jul", albums)
}

// endoint pour la page de la musique "Bolide allemand"
func SdmTrackHandler(w http.ResponseWriter, r *http.Request) {

	//requete vers api se spotify pour obtenir
	//les infos de la musique "bolide allemand"
	//---->> ID Du titre "bolide allemand " = 0EzNyXyU7gHzj2TN8qYThj
	url := "https://api.spotify.com/v1/tracks/0EzNyXyU7gHzj2TN8qYThj"

	body, err := MakeApiRequest(url)

	if err != nil {
		log.Fatal(err)
		fmt.Println("ERREUR LORS DE LA RECUPERATION DES ALBUMS DE JUL")
	}
	//Analyse des reponses json
	var track manager.Tracks
	err = json.Unmarshal(body, &track)

	if err != nil {
		log.Fatal(err)
	}

	// Réponse : Charger et exécuter le template avec les données du titre
	initTemplate.Temp.ExecuteTemplate(w, "sdm", track)
}
