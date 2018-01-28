# Twitter Meme Bot

This is a twitter bot made in [Golang](https://golang.org), It's goal is to look up images on [Reddit](https://reddit.com) and posts them on [Twitter](https://twitter.com)

## Features

- Light weight
- Easily customizable
- Duplicate Image hash checking

## Installation

I use [dep](https://github.com/golang/dep) to manage my dependencies so make sure you have that installed before proceeding.

- Create a twitter application at [apps.twitter.com](https://apps.twitter.com/app/new)
- And create a Reddit app at [reddit.com](https://www.reddit.com/prefs/apps)
- After installing dep copy the contents of [.env.example](.env.example) into your own .env and fill it out with your credentials.
- Now that your credentials are in place you can execute the 2 commands below and your bot should be up and running.
```sh
$ dep ensure
$ go run main.go
```
- you may be prompted to install db drivers like mysql and postgres. just run go get 'required_package' and you're good to go.

You can also run this bot on Heroku, just fill in your `env` info in the settings

## License
[WTFPL License](LICENSE)
