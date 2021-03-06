package main

import (
	"fmt"
	"net/http"

	"github.com/FlaviaBastos/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/dogo":           "https://www.instagram.com/explore/tags/dogoftheday/",
		"/yt":             "https://youtube.com",
		"/myt":            "https://music.youtube.com",
		"/pup":            "https://www.instagram.com/explore/tags/puppylove/",
		"/iggy":           "https://www.instagram.com/explore/tags/italiangreyhound/",
		"/vira":           "https://www.instagram.com/explore/tags/viralata/",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution
    - path: /dogo
      url: https://www.instagram.com/explore/tags/dogoftheday/
    - path: /yt
      url: https://youtube.com
    - path: /myt            
      url: https://music.youtube.com
    - path: /pup
      url: https://www.instagram.com/explore/tags/puppylove/
    - path: /iggy
      url: https://www.instagram.com/explore/tags/italiangreyhound/
    - path: /vira
      url: https://www.instagram.com/explore/tags/viralata/
    `
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
	// http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
