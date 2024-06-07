
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
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
	defer video.Close();
	// headers,
	w.Header().Set("Content-Type","video/mp4");

	io.Copy(w, video);
}

func main() {
	r := mux.NewRouter();
	r.HandleFunc("/", serveVideo);

	err := http.ListenAndServe(":5555",r); if err != nil {
		log.Fatal(err);
	}
}

