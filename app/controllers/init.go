package controllers

import (
	"crypto/md5"
	"github.com/revel/revel"
	"gogo/app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func init() {
	//revel.OnAppStart(Migrate)
	revel.InterceptMethod((*BaseController).Startup, revel.BEFORE)
	revel.InterceptMethod((*BaseController).СheckLoggedIn, revel.BEFORE)
	revel.InterceptMethod((*BaseController).СheckPermissions, revel.BEFORE)
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

	roles, permissions, users := models.Migrations()
	for i := range roles {
		err := connection.DB(revel.Config.StringDefault("mgo.database", "")).C("roles").Insert(models.Role{
			Id:   bson.NewObjectId().Hex(),
			Name: roles[i]})
		if err != nil {
			panic(err)
		}
	}

	for i := range users {
		var role models.Role
		err := connection.DB(revel.Config.StringDefault("mgo.database", "")).C("roles").Find(bson.M{"name": users[i][4]}).One(&role)
		if err != nil {
			panic(err)
		}

		cryptedPassword := md5.Sum([]byte(users[i][3]))

		err = connection.DB(revel.Config.StringDefault("mgo.database", "")).C("users").Insert(models.User{
			Id:        bson.NewObjectId().Hex(),
			Email:     users[i][0],
			FirstName: users[i][1],
			LastName:  users[i][2],
			Password:  string(cryptedPassword[:]),
			Role:      role,
			Joined:    time.Now(),
			Updated:   time.Now()})
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
