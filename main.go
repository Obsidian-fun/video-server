
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
	w.Header().Set("Content-Disposition","attachment; filename=Secret.mp4");

	var buf = make([]byte, 32);
	for i:=0; ;i++ {
		n, err := video.Read(buf); if err == nil {
			fmt.Printf("%s",string(buf[:n])+ "\n");
		}
	}
	if err != nil || err != io.EOF {
		fmt.Println("\n\nerror reading from file\n");
		return
	}
}

func main() {
	r := mux.NewRouter();
	r.Handle("/", http.FileServer(http.Dir(".")));
	r.HandleFunc("/hello", serveVideo);

	err := http.ListenAndServe(":5555",r); if err != nil {
		log.Fatal(err);
	}
}

