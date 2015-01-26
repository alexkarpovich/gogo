package controllers

import (
	"github.com/revel/revel"
	//"gogo/app/models"
	//"gopkg.in/mgo.v2"
	//"os"
)

type User struct {
	*revel.Controller
}

func (this User) Signup() revel.Result {
	return this.Render()
}

func (this User) Login() revel.Result {
	return this.Render()
}

func (this User) Logout() revel.Result {
	return this.Redirect("/admin/user/list")
}