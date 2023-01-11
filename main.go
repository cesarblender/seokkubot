package main

import (
	"net/http"
	"time"

	"mybot/lib"
	"mybot/settings"
)

func main() {
	settings.LoadEnv()

	// hello world http route
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	go http.ListenAndServe(":"+settings.GetEnv("PORT"), nil)

	publishPost()

	go postCommentsReviewLoop()
	time.Sleep(time.Duration(1<<63 - 1))
}

// This function will be called each 30 minutes for every time
func postCommentsReviewLoop() {
	for range time.Tick(time.Minute * 30) {
		publishPost()
	}
}

func publishPost() {
	comments_on_latest_post := lib.GetLastPostValidComments()
	comments_on_latest_post = lib.GetSortedComments(comments_on_latest_post)

	comment_winner := comments_on_latest_post[0]

	lib.CreatePost(
		comment_winner.Message,
		comment_winner.LikeCount,
		len(comments_on_latest_post)-1,
		comment_winner.Attachment.Media.Image.Src,
		comment_winner.CreatedTime,
	)
}
