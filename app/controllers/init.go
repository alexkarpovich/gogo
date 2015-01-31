package controllers

import (
	"github.com/revel/revel"
	"gogo/app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	revel.OnAppStart(Migrate)
	revel.InterceptMethod((*BaseController).Startup, revel.BEFORE)
	revel.InterceptMethod((*BaseController).CheckLoggedIn, revel.BEFORE)
	revel.InterceptMethod((*BaseController).CheckPermissions, revel.BEFORE)
	revel.InterceptMethod((*BaseController).Shutdown, revel.AFTER)
}

func Connect() *mgo.Session {
	HOST := revel.Config.StringDefault("mgo.host", "")
	connection, err := mgo.Dial("mongodb://" + HOST)
	if err != nil {
		panic(err)
	}

	return connection
}

func Migrate() {
	connection := Connect()
	defer connection.Close()

	roles, permissions := models.Migrations()
	for i := range roles {
		err := connection.DB(revel.Config.StringDefault("mgo.database", "")).C("roles").Insert(models.Role{
			Id:   bson.NewObjectId().Hex(),
			Name: roles[i]})
		if err != nil {
			panic(err)
		}
	}

	for controller, action_roles := range permissions {
		for action, roles := range action_roles {
			Roles := make([]models.Role, 0)
			for i := range roles {
				var role models.Role
				err := connection.DB(revel.Config.StringDefault("mgo.database", "")).C("roles").Find(bson.M{"name": roles[i]}).One(&role)
				if err != nil {
					panic(err)
				}
				Roles = append(Roles, role)
			}
			err := connection.DB(revel.Config.StringDefault("mgo.database", "")).C("permissions").Insert(models.Permission{
				Id: bson.NewObjectId().Hex(),
				Resource: models.Resource{
					Controller: controller,
					Action:     action},
				Roles: Roles})
			if err != nil {
				panic(err)
			}
		}
	}
}
