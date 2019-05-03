package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/gophercises/prob_5/models"
)

type link struct {
	URL string `xml:"loc"`
}

type xmlMap struct {
	XMLName   xml.Name `xml:"urlset"`
	NameSpace string   `xml:"xmlns,attr"`
	Links     []link   `xml:"url"`
}

func main() {
	links := make(map[string]int)
	baseURL, path := "https://gophercises.com", "/demos/cyoa"
	models.CreateMap(baseURL, path, 0, 5, links)
	linkArray := make([]link, 0)
	for url := range links {
		linkArray = append(linkArray, link{URL: url})
	}
	xmlMp := xmlMap{NameSpace: baseURL, Links: linkArray}
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent(" ", "	")
	if err := enc.Encode(&xmlMp); err != nil {
		fmt.Println(err)
	}
}
