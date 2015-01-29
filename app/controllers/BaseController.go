package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BaseController struct {
	*revel.Controller
	Storage *mgo.Session
}

func (this *BaseController) Startup() revel.Result {
	HOST := revel.Config.StringDefault("mgo.host","")

	connection,err := mgo.Dial("mongodb://"+HOST)
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

func (this BaseController) Insert(collection string, entity interface {}) {
	err := this.Storage.DB("blog").C(collection).Insert(entity)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) FindAll(collection string, entities ...interface {}) {
	err := this.Storage.DB("blog").C(collection).Find(bson.M{}).All(&entities)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) FindOne(collection string, criteria interface{}, entity interface {}) {
	err := this.Storage.DB("blog").C(collection).Find(criteria).One(entity)
	if err != nil {
		panic(err)
	}
}
