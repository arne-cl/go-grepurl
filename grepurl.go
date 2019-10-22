package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

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

func printAllURLs(t *html.Tokenizer) {
	for {
		tokenType := t.Next()

		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := t.Err()
			if err == io.EOF {
				break //end of file, break out of the loop
			}

			//malformed HTML is "normal", so we'll just continue
			fmt.Fprintf(os.Stderr, "error tokenizing HTML: %v", t.Err())
			continue
		}

		token := t.Token()
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

// htmlTokenizerFromFilePath opens a local HTML file
// and returns a tokenizer of it.
func htmlTokenizerFromFilePath(path string) *html.Tokenizer {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't read file '%s': %v\n", path, err)
	}
	return html.NewTokenizer(file)
}

// htmlTokenizerFromURL takes the URL of an HTML file, downloads it and
// returns a tokenizer of it.
func htmlTokenizerFromURL(url string) *html.Tokenizer {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("error fetching URL: %v\n", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("response status code was %d\n", resp.StatusCode)
	}

	ctype := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ctype, "text/html") {
		log.Fatalf("response content type was %s not text/html\n", ctype)
	}

	//make sure the response body gets closed
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if nil != err {
		fmt.Println("Error reading response body: ", err.Error())
	}

	r := bytes.NewReader(body)
	return html.NewTokenizer(r)
}

// isValidURL tests a string to determine if it is a well-structured URL.
// cf. https://stackoverflow.com/questions/31480710/validate-url-with-standard-package-in-go
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// isValidFilePath checks if a file exists and is not a directory.
func isValidFilePath(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	path := os.Args[1]

	var tokenizer *html.Tokenizer
	if isValidURL(path) {
		tokenizer = htmlTokenizerFromURL(path)
	} else if isValidFilePath(path) {
		tokenizer = htmlTokenizerFromFilePath(path)
	} else {
		log.Fatalf("Path '%v' is not a valid URL / file path.", path)
	}

	printAllURLs(tokenizer)
}
