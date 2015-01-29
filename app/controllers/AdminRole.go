package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"gogo/app/models"
	"gogo/common/db"
	"os"
)

type AdminRole struct {
	BaseController
}

func (this AdminRole) List() revel.Result {

	session,err := db.Connect()
	if err != nil {
		os.Exit(1)
	}

	defer session.Close()

	var roles []models.Role

	err = session.DB("blog").C("roles").Find(bson.M{}).All(&roles)

	if err!=nil {
		os.Exit(1)
	}

	this.RenderArgs["roles"] = roles

	return this.Render()
}

func (this AdminRole) Create(role *models.Role) revel.Result {
	if this.Request.Method == "POST" {

		role.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect(AdminRole.Create)
		}

		session,err := db.Connect()
		if err != nil {
			os.Exit(1)
		}

		defer session.Close()

		err = session.DB("blog").C("roles").Insert(models.Role{
			Id: bson.NewObjectId().Hex(),
			Name: role.Name})

		if err!=nil {
			os.Exit(1)
		}

		return this.Redirect(AdminRole.List)
	} 

	return this.Render()
}

func (this AdminRole) Update(id string) revel.Result {

	session,err := db.Connect()
	if err != nil {
		os.Exit(1)
	}

	defer session.Close()

	if this.Request.Method == "POST" {

		var role *models.Role

		this.Params.Bind(&role, "role")

		role.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect("/admin/role/update/"+id)
		}		

		err = session.DB("blog").C("roles").Update(bson.M{"_id": id}, bson.M{
			"$set": bson.M{
				"name": role.Name}})

		if err != nil {
			os.Exit(1)
		}

		return this.Redirect(AdminRole.List)
	} 	

	var role *models.Role

	err = session.DB("blog").C("roles").Find(bson.M{"_id":id}).One(&role)

	if err!=nil {
		os.Exit(1)
	}

	this.RenderArgs["role"] = role

	return this.Render()
}

func (this AdminRole) Delete(id string) revel.Result {
	session,err := db.Connect()
	if err != nil {
		os.Exit(1)
	}

	defer session.Close()

	err = session.DB("blog").C("roles").Remove(bson.M{"_id":id})

	if err!=nil {
		os.Exit(1)
	}

	return this.Redirect(AdminRole.List)
}
