package reddit

import (
	"net/http"
	"os"
	"encoding/json"
	"log"
	"fmt"
	"strconv"
	"io/ioutil"
	"twitter-meme-bot/database"
	"twitter-meme-bot/structs"
)

func GetThreads() (*[]structs.Thread) {
	println("Checking for new threads to post")

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.reddit.com/r/" + os.Getenv("SUB_REDDIT") + "/controversial/", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Meme-House")

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if response.StatusCode != http.StatusOK || err != nil {
		println("Error getting api data, status code: " + strconv.Itoa(response.StatusCode))
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	s, err := parseResponse([]byte(body))
	if err == nil {
		threads := s.Data.Threads

		filteredThreads := filterThreads(threads)

		defer response.Body.Close()

		return &filteredThreads
	}
	return &[]structs.Thread{}
}

func filterThreads(threads structs.Threads) ([]structs.Thread) {
	FormattedThreads := []structs.Thread{}

	for _, element := range threads.Children {
		if len(element.Data.Preview.Images) > 0 {
			imageUrl := element.Data.Preview.Images[0].Source.URL

			thread := structs.Thread{
				ImageUrl: imageUrl,
				Title:    element.Data.Title,
				Id:       element.Data.ID,
				Author:   element.Data.Author,
			}

			if database.GetThreadById(element.Data.ID) == false {
				FormattedThreads = append(FormattedThreads, thread)
				database.InsertThread(thread)
			}
		}
	}
	return FormattedThreads
}

func parseResponse(body []byte) (*structs.RedditResponse, error) {
	var s = new(structs.RedditResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}