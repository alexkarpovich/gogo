package controllers

import (
	"github.com/revel/revel"
	"gogo/app/models"
	"gopkg.in/mgo.v2/bson"
	"gogo/common/db"
	"crypto/md5"
	"time"
	"os"
)

type AccountUser struct {
	*revel.Controller
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

		connection,err := db.Connect()
		if err != nil {
			os.Exit(1)
		}
		defer connection.Close()

		cryptedPassword := md5.Sum([]byte(user.Password))

		err = connection.DB("blog").C("users").Insert(models.User{
			Id: bson.NewObjectId().Hex(),
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Password: string(cryptedPassword[:]),
			Joined: time.Now(),
			Updated: time.Now()})
		if err != nil {
			os.Exit(1)
		}

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

		connection,err := db.Connect()
		if err != nil {
			os.Exit(1)
		}
		defer connection.Close()

		var loggedInUser models.User

		cryptedPassword := md5.Sum([]byte(user.Password))

		err = connection.DB("blog").C("users").Find(bson.M{"email":user.Email, "password":string(cryptedPassword[:])}).One(&loggedInUser)
		if err != nil {
			os.Exit(1)
		}

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