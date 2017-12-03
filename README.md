go-scrapealerts
=======

# go-scrapealerts - Parses alerts from omnilert and writes to csv file

Written by: Brennan Turner (@BLTSEC)
Website: https://bltsec.ninja

<img width="40%" src=https://raw.githubusercontent.com/BLTSEC/go-webscraper/master/govatar.png></img>


## Requirements
* git
* go
* https://github.com/PuerkitoBio/goquery

## Installation on macOS
	brew install git
	brew install go
	go get github.com/PuerkitoBio/goquery
	
## Installation on Windows
	choco install git.install
	choco install golang
	go get github.com/PuerkitoBio/goquery
	
## Usage
### Build Binary 
	GOOS=linux go build -ldflags="-s -w" alert_scraper.go
### Run Binary
	./alert_scraper --url https://widgets.omnilert.net/-11273 --dst /tmp

