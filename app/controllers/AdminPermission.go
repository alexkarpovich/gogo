package controllers

import (
	"github.com/revel/revel"
	"gogo/app/models"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type AdminPermission struct {
	BaseController
}

func (this AdminPermission) List() revel.Result {

	var permissions []models.Permission

	this._FindAll("permissions", &permissions)

	this.RenderArgs["permissions"] = permissions

	return this.Render()
}

func (this AdminPermission) Create() revel.Result {
	if this.Request.Method == "POST" {

		resource := strings.Split(this.Params.Get("resource"),":")
		roles := strings.Split(this.Params.Get("roles"),"; ")

		Roles := make([]models.Role, 0)
		for i := range roles {
			var role models.Role
			this._FindOne("roles", bson.M{"name":roles[i]}, &role)
			Roles = append(Roles, role)
		}

		permission := models.Permission {
			Id: bson.NewObjectId().Hex(),
			Resource: models.Resource {
				Controller: resource[0],
				Action: resource[1]},
			Roles: Roles}

		permission.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect(AdminPermission.Create)
		}

		this._Insert("permissions", permission)

		return this.Redirect(AdminPermission.List)
	}

	return this.Render()
}

func (this AdminPermission) Update(id string) revel.Result {

	if this.Request.Method == "POST" {

		resource := strings.Split(this.Params.Get("resource"),":")
		roles := strings.Split(this.Params.Get("roles"),"; ")

		Roles := make([]models.Role, 0)
		for i := range roles {
			var role models.Role
			this._FindOne("roles", bson.M{"name":roles[i]}, &role)
			Roles = append(Roles, role)
		}

		permission := models.Permission {
			Id: id,
			Resource: models.Resource {
				Controller: resource[0],
				Action: resource[1]},
			Roles: Roles}

		permission.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect("/admin/permission/update/"+id)
		}

		this._Update("permissions", bson.M{"_id":id}, bson.M{"$set": permission})

		return this.Redirect(AdminPermission.List)
	}

	var permission *models.Permission

	this._FindOne("permissions", bson.M{"_id": id}, &permission)

	this.RenderArgs["permission"] = permission

	return this.Render()
}

func (this AdminPermission) Delete(id string) revel.Result {
	this._Delete("permissions", bson.M{"_id": id})

	return this.Redirect(AdminPermission.List)
}
