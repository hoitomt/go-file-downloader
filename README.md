# File Downloader

1. This downloader requires a configuration file.
1. From the root directory, run `go get` to download dependencies
1. Run the program `go run *.go -c config.yml`. Assumes a configuration file `config.yml` in the project root directory

## Helpful Go add-ins

1. [GoSublime](https://github.com/DisposaBoy/GoSublime)
1. [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)

## ToDo
1. Use Go's built-in utilities to build the request URL
1. Write the file using the downloaded file name. Right now it uses `test.json`
