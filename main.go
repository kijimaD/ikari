package main

import (
	"bytes"
	"log"
	"os"
	"strings"

	ikari "github.com/kijimaD/ikari/lib"
	"golang.org/x/net/html"
)

func main() {
	bs, err := os.ReadFile("input.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(strings.NewReader(string(bs)))
	if err != nil {
		log.Fatal(err)
	}

	ikari.WrapTextWithAnchorRecursive(doc, "div", "a")
	var b bytes.Buffer
	err = html.Render(&b, doc)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("index.html")

	_, err = f.Write([]byte(b.String()))
	if err != nil {
		log.Fatal(err)
	}
}
