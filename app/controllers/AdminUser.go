package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"gogo/app/models"
	"crypto/md5"
	"time"
)

type AdminUser struct {
	BaseController
}

func (this AdminUser) List() revel.Result {
	var users []models.User

	this._FindAll("users", &users)	

	this.RenderArgs["users"] = users

	return this.Render()
}

func (this AdminUser) Create(user *models.User) revel.Result {
	if this.Request.Method == "POST" {

		user.Validate(this.Validation)

		this.Validation.Required(user.Role)
		this.Validation.Required(user.Password)
		this.Validation.Required(user.ConfirmPassword)
		this.Validation.Required(user.Password==user.ConfirmPassword).Message("Not matching passwords")

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect(AdminUser.Create)
		}

		cryptedPassword := md5.Sum([]byte(user.Password))

		var role models.Role

		this._FindOne("roles", bson.M{"_id":this.Params.Get("role")}, &role)

		this._Insert("users", models.User{
			Id: bson.NewObjectId().Hex(),
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Role: role,
			Password: string(cryptedPassword[:]),
			Joined: time.Now(),
			Updated: time.Now()})

		return this.Redirect(AdminUser.List)
	} 

	var roles []models.Role

	this._FindAll("roles", &roles)

	this.RenderArgs["roles"] = roles

	return this.Render()
}

func (this AdminUser) Update(id string) revel.Result {

	if this.Request.Method == "POST" {

		var user *models.User

		this.Params.Bind(&user, "user")

		user.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect("/admin/user/update/"+id)
		}	

		var role *models.Role	

		this._FindOne("roles", bson.M{"_id":this.Params.Get("role")}, &role)
		this._Update("users", bson.M{"_id": id}, bson.M{
			"$set": bson.M{
				"email": user.Email,
				"firstName": user.FirstName,
				"lastName": user.LastName,
				"role": role,
				"updated": time.Now()}})

		return this.Redirect(AdminUser.List)
	} 	

	var user *models.User
	var roles []*models.Role

	this._FindOne("users", bson.M{"_id":id}, &user)
	
	this._FindAll("roles", &roles)

	this.RenderArgs["user"] = user
	this.RenderArgs["roles"] = roles

	return this.Render()
}

func (this AdminUser) Delete(id string) revel.Result {
	this._Delete("users", bson.M{"_id":id})

	return this.Redirect(AdminUser.List)
}

