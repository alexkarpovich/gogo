package controllers

import (
	"github.com/revel/revel"
	"gogo/app/models"
	"gopkg.in/mgo.v2/bson"
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

func (this AdminPermission) Create(permission *models.Permission) revel.Result {
	if this.Request.Method == "POST" {

		permission.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect(AdminPermission.Create)
		}

		this._Insert("permissions", models.Permission{
			Id:       bson.NewObjectId().Hex(),
			Resource: permission.Resource,
			Roles:    permission.Roles})

		return this.Redirect(AdminPermission.List)
	}

	return this.Render()
}

func (this AdminPermission) Update(id string) revel.Result {

	if this.Request.Method == "POST" {

		var permission *models.Permission

		this.Params.Bind(&permission, "permission")

		permission.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect("/admin/permission/update/" + id)
		}

		this._Update("permissions", bson.M{"_id": id}, bson.M{
			"$set": bson.M{
				"resource": bson.M{
					"controller": "",
					"action":     ""},
				"roles": bson.M{}}})

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
