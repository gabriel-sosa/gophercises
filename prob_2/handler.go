package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

//MapHandler will handle the map
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

//YAMLHandler will handle the yaml
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var paths []pathURL
	err := yaml.Unmarshal(yml, &paths)
	if err != nil {
		return nil, err
	}
	pathsToURL := make(map[string]string)
	for _, value := range paths {
		pathsToURL[value.Path] = value.URL
	}
	return MapHandler(pathsToURL, fallback), nil
}

type pathURL struct {
	Path string `yalm:"path"`
	URL  string `yalm:"url"`
}
