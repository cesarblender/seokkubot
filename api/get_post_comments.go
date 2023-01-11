package api

import (
	"encoding/json"
	"log"

	fb "github.com/huandu/facebook/v2"

	"mybot/settings"
)

type Image struct {
	Src string `json:"src"`
}

type Media struct {
	Image Image `json:"image"`
}

type Attachment struct {
	Media Media `json:"media"`
}

type Comment struct {
	Attachment  Attachment `json:"attachment"`
	ID          string     `json:"id"`
	Message     string     `json:"message"`
	CreatedTime string     `json:"created_time"`
	LikeCount   int        `json:"like_count"`
}

type FacebookResponse struct {
	Data []Comment `json:"data"`
}

func GetPostComments(PostID string) []Comment {
	endpoint := "/" + PostID + "/comments"

	res, err := fb.Get(endpoint, fb.Params{
		"fields":       "message,id,attachment,like_count,created_time",
		"access_token": settings.GetEnv("ACCESS_TOKEN"),
	})
	if err != nil {
		log.Fatal(err)
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	var myresponse FacebookResponse
	if err := json.Unmarshal(jsonStr, &myresponse); err != nil {
		log.Fatal(err)
	}

	data := myresponse.Data

	if len(data) == 0 {
		return []Comment{SendRandomCatImage()}
	}

	return data
}

func SendRandomCatImage() Comment {
	message := "Crossover con Random Cat - Bot.\nNo hay publicaciones pendientes en los comentarios. Comenta '!publicame', seguido de una imagen para comenzar a publicar contenido."

	return Comment{
		Message: message,
		ID:      "",
		Attachment: Attachment{
			Media{
				Image: Image{
					Src: GetRandomCatImage(),
				},
			},
		},
		CreatedTime: "#########",
		LikeCount:   -1,
	}
}
