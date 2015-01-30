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

func (this BaseController) _Insert(collection string, entity interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Insert(entity)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) _FindAll(collection string, entities interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Find(bson.M{}).All(entities)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) _FindOne(collection string, criterion interface{}, entity interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Find(criterion).One(entity)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) _Update(collection string, criterion interface{}, entity interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Update(criterion, entity)
	if err != nil {
		panic(err)
	}
}

func (this BaseController) _Delete(collection string, criterion interface{}) {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Remove(criterion)
	if err != nil {
		panic(err)
	}
}

func (this *BaseController) CheckLoggedIn() revel.Result {
	if id, ok := this.Session["user"]; ok {

		var loggedInUser *models.User

		this._FindOne("users", bson.M{"_id": id}, &loggedInUser)

		this.RenderArgs["loggedInUser"] = loggedInUser

		return nil
	}

	delete(this.RenderArgs, "loggedInUser")

	return nil
}

func (this *BaseController) CheckAccess() revel.Result {
	permissions := models.GetPermissions()
	if user, ok := this.RenderArgs["loggedInUser"]; ok {
		if actions, ok := permissions[user.(*models.User).Role.Name][this.Name]; ok {
			for _, action := range actions {
		        if action == this.MethodName {
		            return nil
		        }
		    }
		}
		return this.Redirect("/")
	} 

	if actions, ok := permissions["Guest"][this.Name]; ok {
			for _, action := range actions {
		        if action == this.MethodName {
		            return nil
		        }
		    }
		}
	
	return this.Redirect("/")	
}