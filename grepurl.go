package main

import (
	//"io/ioutil"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func printAnchorURL(token html.Token) {
	for _, attrib := range token.Attr {
		if attrib.Key == "href" {
			fmt.Printf("%v\n", attrib.Val)
		}
	}
}

func printImageURL(token html.Token) {
	for _, attrib := range token.Attr {
		if attrib.Key == "src" {
			fmt.Printf("%v\n", attrib.Val)
		}
	}
}

func main() {
	// TODO: check if it is a file
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file '%s': %v\n", filepath, err)
		os.Exit(1)
	}

	tokenizer := html.NewTokenizer(file)

	for {
		tokenType := tokenizer.Next()

		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break //end of file, break out of the loop
			}

			//malformed HTML is "normal", so we'll just continue
			fmt.Fprintf(os.Stderr, "error tokenizing HTML: %v", tokenizer.Err())
			continue
		}

		token := tokenizer.Token()
		if tokenType == html.StartTagToken {
			if token.DataAtom == atom.A { // <a> element
				printAnchorURL(token)
			}
		} else if tokenType == html.SelfClosingTagToken {
			if token.DataAtom == atom.Img { // <img /> element
				printImageURL(token)
			}
		}
	}
}
