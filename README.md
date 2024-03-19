# go-grepurl

`grepurl` is a command-line tool written in Go that extracts URLs from HTML content.
It can process HTML from both local files and remote web pages,
outputting all found URLs to the console.


## Features

- Extract URLs from local HTML files.
- Fetch HTML from remote URLs and extract links.
- Support for both `<a href="...">` and `<img src="...">` tags.
- Simple and easy-to-use command-line interface.

## Installation

### From Source

To install `grepurl` from source, ensure you have Go installed on your system
([official Go installation guide](https://golang.org/doc/install)).
hen follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/arne-cl/go-grepurl.git
cd go-grepurl
```

2. Build the binary:

```bash
go build -o grepurl
```

3. (Optional) Move the binary to a location in your PATH for global access:

```bash
sudo mv grepurl /usr/local/bin/
```

## Usage

After installation, you can use `grepurl` by providing a local file path or a remote URL as an argument.
Here are some examples:

- Extract URLs from a local HTML file:

```bash
grepurl /path/to/your/file.html
```

- Extract URLs from a webpage:

```bash
grepurl https://example.com
```

`grepurl` will output all found URLs to the console.


## License

`go-grepurl` is released under the MIT License.
