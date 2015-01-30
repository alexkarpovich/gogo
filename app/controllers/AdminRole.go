package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"gogo/app/models"
)

type AdminRole struct {
	BaseController
}

func (this AdminRole) List() revel.Result {

	var roles []models.Role

	this._FindAll("roles", &roles)

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

		this._Insert("roles", models.Role{
			Id: bson.NewObjectId().Hex(),
			Name: role.Name})

		return this.Redirect(AdminRole.List)
	} 

	return this.Render()
}

func (this AdminRole) Update(id string) revel.Result {

	if this.Request.Method == "POST" {

		var role *models.Role

		this.Params.Bind(&role, "role")

		role.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect("/admin/role/update/"+id)
		}		

		this._Update("roles", bson.M{"_id": id}, bson.M{
			"$set": bson.M{
				"name": role.Name}})

		return this.Redirect(AdminRole.List)
	} 	

	var role *models.Role

	this._FindOne("roles", bson.M{"_id":id}, &role)

	this.RenderArgs["role"] = role

	return this.Render()
}

func (this AdminRole) Delete(id string) revel.Result {
	this._Delete("roles", bson.M{"_id":id})

	return this.Redirect(AdminRole.List)
}
