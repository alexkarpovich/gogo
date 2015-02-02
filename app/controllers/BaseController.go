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
	this.Storage = Connect()

	return nil
}

func (this *BaseController) Shutdown() revel.Result {
	if this.Storage != nil {
		this.Storage.Close()
	}

	return nil
}

func (this BaseController) _Insert(collection string, entity interface{}) error {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Insert(entity)
	return err
}

func (this BaseController) _FindAll(collection string, entities interface{}) error {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Find(bson.M{}).All(entities)
	return err
}

func (this BaseController) _FindOne(collection string, criterion interface{}, entity interface{}) error {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Find(criterion).One(entity)
	return err
}

func (this BaseController) _Update(collection string, criterion interface{}, entity interface{}) error {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Update(criterion, entity)
	return err
}

func (this BaseController) _Delete(collection string, criterion interface{}) error {
	err := this.Storage.DB(revel.Config.StringDefault("mgo.database", "")).C(collection).Remove(criterion)
	return err
}

func (this *BaseController) СheckLoggedIn() revel.Result {
	if id, ok := this.Session["user"]; ok {

		var loggedInUser *models.User

		this._FindOne("users", bson.M{"_id": id}, &loggedInUser)

		this.RenderArgs["loggedInUser"] = loggedInUser

		return nil
	}

	delete(this.RenderArgs, "loggedInUser")

	return nil
}

func (this *BaseController) СheckPermissions() revel.Result {
	var currentUserRole string = "Guest"
	if user, ok := this.RenderArgs["loggedInUser"]; ok {
		currentUserRole = user.(*models.User).Role.Name
	}

	var permission models.Permission

	err := this._FindOne("permissions", bson.M{
		"resource": bson.M{
			"controller": this.Name,
			"action":     this.MethodName},
		"roles": bson.M{
			"$elemMatch": bson.M{
				"name": currentUserRole}}}, &permission)
	if err != nil {
		this.Storage.Close()
		return this.Redirect("/")
	}
	return nil
}
