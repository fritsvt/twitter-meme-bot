# Twitter Meme Bot

This is a twitter bot made in [Golang](https://golang.org), It's goal is to look up images on [Reddit](https://reddit.com) and posts them on [Twitter](https://twitter.com)

## Features

- Light weight
- Easily customizable
- Duplicate Image hash checking
- Tweet scheduling

## Usage

In general there isn't much you have to do to run this bot besides configuring it once.
But there is a Tweet scheduling feature and if you want to access that you can do so on the `/schedule` route

## Installation

I use [dep](https://github.com/golang/dep) to manage my dependencies so make sure you have that installed before proceeding.

- Create a twitter application at [apps.twitter.com](https://apps.twitter.com/app/new)
- And create a Reddit app at [reddit.com](https://www.reddit.com/prefs/apps)
- After installing dep copy the contents of [.env.example](.env.example) into your own .env and fill it out with your credentials.
- Now that your credentials are in place you can execute the 2 commands below and your bot should be up and running.
```sh
$ go build main.go
$ ./main
```
- you may be prompted to install db drivers like mysql and postgres. just run go get 'required_package' and you're good to go.

## Dependencies

The bot uses the following dependencies:
- [github.com/jinzhu/gorm](https://github.com/jinzhu/gorm)
- [github.com/turnage/graw](https://github.com/turnage/graw)
- [github.com/ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda)
- [github.com/devedge/imagehash](https://github.com/devedge/imagehash)
- [github.com/gorilla/mux](https://github.com/gorilla/mux)
- [github.com/joho/godotenv](https://github.com/joho/godotenv)

## License
[WTFPL License](LICENSE)
