package api

import (
	"log"

	fb "github.com/huandu/facebook/v2"

	"mybot/settings"
)

func GetLastPost() string {
	res, err := fb.Get("/me/feed", fb.Params{
		"fields":       "id",
		"access_token": settings.GetEnv("ACCESS_TOKEN"),
	})

	if err != nil {
		log.Fatal(err)
	}

	data := res.GetField("data").([]interface{})

	lastpost := data[0]

	return lastpost.(map[string]interface{})["id"].(string)
}
