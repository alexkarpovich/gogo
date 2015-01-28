package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"gogo/app/models"
	"gogo/common/db"
	"crypto/md5"
	"time"
	"fmt"
	"os"
)

type AdminUser struct {
	*revel.Controller
}

func (this AdminUser) List() revel.Result {
	session,err := db.Connect()
	if err != nil {
		fmt.Printf("Error connection")
		os.Exit(1)
	}

	defer session.Close()

	var users []models.User

	err = session.DB("blog").C("users").Find(bson.M{}).All(&users)

	if err!=nil {
		fmt.Printf("Find error")
		os.Exit(1)
	}

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

		session,err := db.Connect()
		if err != nil {
			os.Exit(1)
		}

		defer session.Close()

		cryptedPassword := md5.Sum([]byte(user.Password))

		var role models.Role

		err = session.DB("blog").C("roles").Find(bson.M{"_id":this.Params.Get("role")}).One(&role)

		err = session.DB("blog").C("users").Insert(models.User{
			Id: bson.NewObjectId().Hex(),
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Role: role,
			Password: string(cryptedPassword[:]),
			Joined: time.Now(),
			Updated: time.Now()})

		if err!=nil {
			os.Exit(1)
		}

		return this.Redirect(AdminUser.List)
	} 

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

func (this AdminUser) Update(id string) revel.Result {

	session,err := db.Connect()
	if err != nil {
		fmt.Printf("Error connection")
		os.Exit(1)
	}

	defer session.Close()

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

		err = session.DB("blog").C("roles").Find(bson.M{"_id":this.Params.Get("role")}).One(&role)

		if err!=nil {
			os.Exit(1)
		}

		err = session.DB("blog").C("users").Update(bson.M{"_id": id}, bson.M{
			"$set": bson.M{
				"email": user.Email,
				"firstName": user.FirstName,
				"lastName": user.LastName,
				"role": role,
				"updated": time.Now()}})

		if err != nil {
			os.Exit(1)
		}

		return this.Redirect(AdminUser.List)
	} 	

	var user *models.User
	var roles []*models.Role

	err = session.DB("blog").C("users").Find(bson.M{"_id":id}).One(&user)

	if err!=nil {
		os.Exit(1)
	}

	err = session.DB("blog").C("roles").Find(bson.M{}).All(&roles)

	if err!=nil {
		os.Exit(1)
	}

	this.RenderArgs["user"] = user
	this.RenderArgs["roles"] = roles

	return this.Render()
}

func (this AdminUser) Delete(id string) revel.Result {
	session,err := db.Connect()
	if err != nil {
		os.Exit(1)
	}

	defer session.Close()

	err = session.DB("blog").C("users").Remove(bson.M{"_id":id})

	if err!=nil {
		fmt.Printf("Delete error")
		os.Exit(1)
	}

	return this.Redirect(AdminUser.List)
}

