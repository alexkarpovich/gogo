package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BaseController struct {
	*revel.Controller
	Storage *mgo.Database
}

func (this *BaseController) Startup() revel.Result {
	HOST := revel.Config.StringDefault("mgo.host","")
	DATABASE := revel.Config.StringDefault("mgo.database","")

	connection,err := mgo.Dial("mongodb://"+HOST)
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	this.Storage = connection.DB(DATABASE)

	return nil
}

func (this *BaseController) Shutdown() revel.Result {
	if this.Storage != nil {
		this.Storage.Logout()
	}

	return nil
}

func Insert(this *BaseController, collection string, entity interface {}) {
	err := this.Storage.C(collection).Insert(entity)
	if err != nil {
		panic(err)
	}
}

func FindAll(this *BaseController, collection string, entities []interface {}) {
	err := this.Storage.C(collection).Find(bson.M{}).All(&entities)
	if err != nil {
		panic(err)
	}
}