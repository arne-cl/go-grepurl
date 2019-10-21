package main

import (
	//"io/ioutil"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

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

		// fmt.Printf("%v\n", tokenType)

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if "a" == token.Data { // if this is an <a> element
				for _, attrib := range token.Attr {
					if attrib.Key == "href" {
						fmt.Printf("%v\n", attrib.Val)
					}
				}
			}
		}

	}

}
