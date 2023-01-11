package lib

import (
	"mybot/api"
	"regexp"
	"strings"
)

func GetLastPostValidComments() []api.Comment {
	latest_post := api.GetLastPost()
	data := api.GetPostComments(latest_post)

	var valid_comments []api.Comment

	for i := 0; i < len(data); i++ {
		comment := data[i]

		if comment.Attachment == (api.Attachment{}) {
			continue
		}

		id := comment.ID
		message := comment.Message
		created_time := comment.CreatedTime
		like_count := comment.LikeCount
		img := comment.Attachment.Media.Image.Src

		valid_keyword, _ := regexp.Compile("^!(publicame|publícame|publìcame|publocame|publicamee|piblicame|poblicame)")

		keyword_matchs := valid_keyword.FindAllString(strings.ToLower(message), 1)
		if len(keyword_matchs) == 0 {
			continue
		}

		message = string(valid_keyword.ReplaceAll([]byte(message), []byte("")))
		message = strings.Replace(message, "\n", "", 1)

		valid_comments = append(valid_comments, api.Comment{
			ID:      id,
			Message: message,
			Attachment: api.Attachment{
				Media: api.Media{
					Image: api.Image{
						Src: img,
					},
				},
			},
			LikeCount:   like_count,
			CreatedTime: created_time,
		})
	}

	if len(valid_comments) == 0 {
		return []api.Comment{api.SendRandomCatImage()}
	}

	return valid_comments
}
