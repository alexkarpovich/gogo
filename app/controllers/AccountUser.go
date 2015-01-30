package controllers

import (
	"github.com/revel/revel"
	"gogo/app/models"
	"gopkg.in/mgo.v2/bson"
	"crypto/md5"
	"time"
)

type AccountUser struct {
	BaseController
}

func (this AccountUser) Signup() revel.Result {

	if this.Request.Method == "POST" {
		
		var user models.User
		this.Params.Bind(&user,"user")

		user.Validate(this.Validation)

		this.Validation.Required(user.Password)
		this.Validation.Required(user.ConfirmPassword)
		this.Validation.Required(user.Password==user.ConfirmPassword).Message("Not matching passwords")

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect(AccountUser.Signup)
		}

		cryptedPassword := md5.Sum([]byte(user.Password))

		var role models.Role

		this._FindOne("roles", bson.M{"name":"Employee"}, &role)

		this._Insert("users", models.User{
			Id: bson.NewObjectId().Hex(),
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Role: role,
			Password: string(cryptedPassword[:]),
			Joined: time.Now(),
			Updated: time.Now()})

		return this.Redirect("/admin/user/list")
	}

	return this.Render()
}

func (this AccountUser) Login() revel.Result {

	if this.Request.Method == "POST" {
		
		var user models.User
		this.Params.Bind(&user,"user")

		this.Validation.Required(user.Email)
		this.Validation.Required(user.Password)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect(AccountUser.Login)
		}

		var loggedInUser models.User

		cryptedPassword := md5.Sum([]byte(user.Password))

		this._FindOne("users", bson.M{"email":user.Email, "password":string(cryptedPassword[:])}, &loggedInUser)

		this.Session = make(revel.Session)

		this.Session["user"] = loggedInUser.Id

		return this.Redirect("/admin/user/list")
	}

	return this.Render()
}

func (this AccountUser) Logout() revel.Result {

	for k := range this.Session {
		delete(this.Session, k)
	}

	return this.Redirect("/")
}

func (this AccountUser) Profile() revel.Result {

	if _, ok := this.RenderArgs["loggedInUser"]; !ok {
		return this.Redirect(AccountUser.Login)
	}

	return this.Render()
}

func (this AccountUser) Retrieve(id string) revel.Result{
	var user *models.User

	this._FindOne("users", bson.M{"_id":id}, &user)

	this.RenderArgs["user"] = user

	return this.Render()
}