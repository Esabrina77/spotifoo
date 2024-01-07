package roots

import (
	"log"
	"net/http"
	"spotifoo/controller"
)

func InitServe() {
	FileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", FileServer))
	http.HandleFunc("/welcome", controller.WelcomeHandler)
	http.HandleFunc("/jul", controller.JulAlbumHandler)
	http.HandleFunc("/sdm", controller.SdmTrackHandler)
	if err := http.ListenAndServe(controller.Port, nil); err != nil {
		log.Fatal(err)
	}
}
