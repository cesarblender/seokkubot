package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RandomCatRequest struct {
	Url string `json:"url"`
}

const RandomCatEndpoint = "https://api.thecatapi.com/v1/images/search"

func GetRandomCatImage() string {
	res, err := http.Get(RandomCatEndpoint)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
	}
	var randomCatRequest []RandomCatRequest
	err = json.Unmarshal([]byte(body), &randomCatRequest)
	if err != nil {
		log.Println(err)
	}
	imgUrl := randomCatRequest[0].Url

	return imgUrl
}
