package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type gopher struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

//HandleRoutes handle the routes
func HandleRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		jsonBlob, err := ioutil.ReadFile("./prob_3/gopher.json")
		check(err)
		story := jsonToMap(jsonBlob)
		if info, ok := story[path]; ok {
			render(w, "index", info)
		} else {
			render(w, "index", story["intro"])
		}
	})
}

func render(w http.ResponseWriter, name string, data interface{}) {
	rawTpl, err := ioutil.ReadFile("./prob_3/views/" + name + ".html")
	check(err)
	temp := string(rawTpl)
	t, err := template.New("webpage").Parse(temp)
	check(err)
	err = t.Execute(w, data)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func jsonToMap(data []byte) (i map[string]gopher) {
	err := json.Unmarshal(data, &i)
	check(err)
	return
}
