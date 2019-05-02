package main

import (
	"io/ioutil"
	"regexp"
)

func main() {
	test, _ := ioutil.ReadFile("./prob_4/ex4.html")
	regExpComment := regexp.MustCompile(`<!--.+-->`)
	test = regExpComment.ReplaceAll(test, []byte(""))
	regExpLink := regexp.MustCompile(`(?i)<a\s*?.*?\s*?href="([^\s]+)"\s*?.*?\s*?>\s*?(<\/?(\s|\S)*?>)*?\s*?([^<>]+)\s*?(<\/?(\s|\S)*?>)*?\s*?</a>`)
	for _, i := range regExpLink.FindAllSubmatch(test, -1) {
		println("url: " + string(i[1]))
		println("content: " + string(i[4]))
	}
}
