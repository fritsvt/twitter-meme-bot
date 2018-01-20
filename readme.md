# Twitter Meme Bot

This is a twitter bot made in [Golang](https://golang.org), It's goal is to look up images on [Reddit](https://reddit.com) and posts them on [Twitter](https://twitter.com)
I created the bot to run of the [@memesforfemes](https://twitter.com/memesforfemes) Twitter account but feel free to make your own.

## Installation

I use [dep](https://github.com/golang/dep) to manage my dependencies so make sure you have that installed before proceeding.

- Create a twitter application at [apps.twitter.com](https://apps.twitter.com/app/new)
- After installing dep copy the contents of [.env.example](.env.example) into your own .env and fill it out with your credentials.
- Now that your credentials are in place you can execute the 2 commands below and your bot should be up and running.
```sh
$ dep ensure
$ go run main.go
```

You can also run this bot on Heroku, just change your `DB_DRIVER` to postgres and update the optional `DATABASE_URL` with the url you receive from Heroku

## License
[WTFPL License](LICENSE)