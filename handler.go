package urlshort

import (
	"log"
	"net/http"

	"github.com/go-yaml/yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	hfn := func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Path

		if match, ok := pathsToUrls[v]; ok {
			http.Redirect(w, r, match, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
	return http.HandlerFunc(hfn)
}

// Yaml is a slice of path and urls
type Yaml []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var y Yaml
	err := yaml.Unmarshal(yml, &y)
	if err != nil {
		log.Fatalln(err)
	}

	mapper := make(map[string]string)
	for _, v := range y {
		mapper[v.Path] = v.URL
	}
	return MapHandler(mapper, fallback), nil
}
