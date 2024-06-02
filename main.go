
package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func serveVideo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there Friend\n\n");

	cookie := &http.Cookie{
		Name:"Example.com",
		Value:"1234",
		Domain:"happygolucky.com",
	}
	http.SetCookie(w, cookie);

	fmt.Fprintf(w, fmt.Sprintf("Name:%s,\nDomain:%s\n", cookie.Name, cookie.Domain));

	ctx := r.Context();
	fmt.Fprintf(w, ctx);

}

func main() {
	r := mux.NewRouter();
	r.Handle("/", http.FileServer(http.Dir(".")));
	r.HandleFunc("/hello", serveVideo);

	err := http.ListenAndServe(":5555",r); if err != nil {
		log.Fatal(err);
	}
}

