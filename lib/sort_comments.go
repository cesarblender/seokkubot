package lib

import (
	"mybot/api"
	"sort"
)

func GetSortedComments(comments []api.Comment) []api.Comment {
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].LikeCount > comments[j].LikeCount
	})

	return comments
}
