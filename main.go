
package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/gorilla/mux"
)

func serveVideo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving video...");
	video ,err := os.Open("./videos/Secret.mp4");
	if err != nil {
		http.Error(w, "Video not found",http.StatusNotFound);
		return
	}

	// headers,
	w.Header().Set("Content-Type","video/mp4");
	//w.Header().Set("Content-Disposition","attachment; filename=Secret.mp4");

	http.ServeFile(w, r, video);
}

func main() {
	r := mux.NewRouter();
	r.Handle("/", http.FileServer(http.Dir(".")));
	r.HandleFunc("/hello", serveVideo);

	err := http.ListenAndServe(":5555",r); if err != nil {
		log.Fatal(err);
	}
}

