package controllers

import (
	"github.com/revel/revel"
	"gogo/app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BaseController struct {
	*revel.Controller
	Storage *mgo.Session
}

func (this *BaseController) Startup() revel.Result {
	HOST := revel.Config.StringDefault("mgo.host", "")

	connection, err := mgo.Dial("mongodb://" + HOST)
	if err != nil {
		panic(err)
	}

	this.Storage = connection

	return nil
}

func (this *BaseController) Shutdown() revel.Result {
	if this.Storage != nil {
		this.Storage.Close()
	}

	return nil
}

func (this BaseController) InsertEntity(collection string, entity interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Insert(entity)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) FindAllEntities(collection string, entities interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Find(bson.M{}).All(entities)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) FindOneEntity(collection string, criteria interface{}, entity interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Find(criteria).One(entity)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) UpdateEntity(collection string, criteria interface{}, entity interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Update(criteria, entity)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) DeleteEntity(collection string, criteria interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Remove(criteria)
	if err != nil {
		panic(err)
	}
}

func (this *BaseController) CheckLoggedIn() revel.Result {
	if id, ok := this.Session["user"]; ok {

		var loggedInUser *models.User

		this.FindOneEntity("users", bson.M{"_id": id}, &loggedInUser)

		this.RenderArgs["loggedInUser"] = loggedInUser

		return nil
	}

	delete(this.RenderArgs, "loggedInUser")

	return nil
}
