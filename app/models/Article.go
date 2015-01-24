package models

import (
	//"gopkg.in/mgo.v2/bson"
	"github.com/revel/revel"
	"time"
)

type Article struct {
	Id 		string `bson:"_id"`
	Title	string `bson:"title"`
	Content string `bson:"content"`
	Author 	string `bson:"author"`
	Created time.Time `bson:"created"`
	Updated time.Time `bson:"updated"`
}

func (article Article) Validate(v *revel.Validation) {
	v.Required(article.Title)
	v.Required(article.Content)
	v.Required(article.Author)
}