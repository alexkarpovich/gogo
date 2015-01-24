package db

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
)

func Connect() (*mgo.Session,error) {
	HOST := revel.Config.StringDefault("mgo.host","")

	return mgo.Dial("mongodb://"+HOST)
}