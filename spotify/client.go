package spotify

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Client struct {
}

func GetClient() *Client {
	return &Client{}
}

func (c *Client) ApiSearchTrackByISCR(iscr string) SearchResponse {
	var data SearchResponse

	baseUrl := "https://api.spotify.com/"
	url := baseUrl + "v1/search?type=track&q=isrc%3A" + iscr

	req, _ := http.NewRequest("GET", url, nil)

	token := os.Getenv("BEARER_TOKEN")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		log.Fatal("renew the token")
		return data
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&data)

	return data
}
