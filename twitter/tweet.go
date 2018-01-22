package twitter

import (
	"os"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"net/http"
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"twitter-meme-bot/structs"
)

var api anaconda.TwitterApi

func Setup() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_PUBLIC"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api = *anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
}

func SendTweet(thread structs.Thread) {
	response, e := http.Get(thread.ImageUrl)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		str := base64.StdEncoding.EncodeToString([]byte(bodyString))

		res, err := api.UploadMedia(str);
		if err != nil {
			log.Fatal(err)
		}

		v := url.Values{}

		v.Set("media_ids", strconv.FormatInt(res.MediaID, 10))

		println("Posting tweet: " + thread.RedditId)
		//api.PostTweet(thread.Title + " #dankmemes",  v);
	}

}