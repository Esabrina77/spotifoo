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

// LE TOKEN EXPIRE APRES CHAQUE 3600sec soit  1H
var token = "BQArPyr4PmJyG4wfajesp373HWfB2eO6TDnzH7fL1MscRWRlARNYX2URe9rflT0ETAlsGIHH1GqzkxP3o1jyP-KrDC6Gwuh00PrKjtvAf6EX-qUfhUM"

// requete Get vers l'api de spotify
// et reception des données obtenues
func MakeApiRequest(url string, token string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("log: \tRequest failed!\n", err)
		manager.PrintColorResult("purple", "log: \tRequest failed!\n")
	}

	//gestion du token (en-tete d'authetification)

	req.Header.Add("Authorization", "Bearer "+token)

	res, errRes := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		log.Fatal("log:\tSending request error!\n", errRes)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal("Reading body request error !\n", errBody)
	}

	return body, nil

}

//endoint pour la page des albums de jul

func JulAlbumHandler(w http.ResponseWriter, r *http.Request) {

	//requete vers api se spotify pour obtenir
	//les infos des albums de jul

	//ID DE Jul = 3IW7ScrzXmPvZhB27hmfgy
	url := "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy/albums"

	body, err := MakeApiRequest(url, token)

	if err != nil {
		log.Fatal(err)
		manager.PrintColorResult("red", "ERREUR LORS DE LA RECUPERATION DES ALBUMS DE JUL")
	}
	//Analyse des reponses json
	var resp manager.AlbumsData
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("makeApiRequest()\tbody: %#v\n", resp)

	// Réponse : Charger et exécuter le template avec les données d'albums
	initTemplate.Temp.ExecuteTemplate(w, "jul", resp)

}

// endoint pour la page de la musique "Bolide allemand"
func SdmTrackHandler(w http.ResponseWriter, r *http.Request) {

	//requete vers api se spotify pour obtenir
	//les infos de la musique "bolide allemand"
	//---->> ID Du titre "bolide allemand " = 0EzNyXyU7gHzj2TN8qYThj
	url := "https://api.spotify.com/v1/tracks/0EzNyXyU7gHzj2TN8qYThj"

	body, err := MakeApiRequest(url, token)

	if err != nil {
		log.Fatal(err)
		manager.PrintColorResult("red", "ERREUR LORS DE LA RECUPERATION DE BOLIDE ALLEMAND")
	}
	//Analyse des reponses json
	var track manager.TrackData
	err = json.Unmarshal(body, &track)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("makeApiRequest()\tbody: %#v\n", track)

	// Réponse : Charger et exécuter le template avec les données du titre
	initTemplate.Temp.ExecuteTemplate(w, "sdm", track)
}
