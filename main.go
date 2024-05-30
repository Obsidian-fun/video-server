
package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {

	const dirName = ".";
	const port = ":5555";


	http.Handle("/",addHeaders(http.FileServer(http.Dir(dirName))));
	fmt.Printf("Serving on localhost%v\n",port);

	err := (http.ListenAndServe(port,nil)); if err != nil {
		log.Fatal(err);
	}
}

// addHeaders is a custom func, to bypass CORS
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin","*");
		h.ServeHTTP(w, r);
	}
}

