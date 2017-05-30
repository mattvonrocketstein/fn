package route

import (
	"gitlab-odx.oracle.com/odx/functions/examples/blog/database"
	"gopkg.in/mgo.v2/bson"
)

func HandlePostList(db *database.Database, auth map[string]interface{}) {
	posts, err := db.GetPosts([]bson.M{})
	if err != nil {
		SendError("Couldn't retrieve posts")
		return
	}

	SendResponse(Response{
		"posts": posts,
	})
}
