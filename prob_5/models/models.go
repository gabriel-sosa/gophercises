package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

//CreateMap uses recursion to create a map with all the links in the url
func CreateMap(baseURL, path string, dept, maxDept int, links map[string]int) {
	url := baseURL + path
	if i, ok := links[path]; ok {
		if dept < i {
			links[path] = dept
		}
	} else {
		links[path] = dept
		template := get(url)
		lks := parseLinks(template)
		if dept < maxDept {
			for _, val := range lks {
				CreateMap(baseURL, val, dept+1, maxDept, links)
			}
		}
	}
}

func get(url string) string {
	resp, err := http.Get(url)
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return string(body)
}

func parseLinks(tpl string) []string {
	links := make([]string, 0)
	regExp := regexp.MustCompile(`<a\s+href="(/.+)"\s*>`)
	for _, val := range regExp.FindAllSubmatch([]byte(tpl), -1) {
		links = append(links, string(val[1]))
	}
	return links
}

func check(i interface{}) {
	if i != nil {
		log.Fatalln(i)
	}
}
