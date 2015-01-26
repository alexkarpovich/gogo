// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tAccountUser struct {}
var AccountUser tAccountUser


func (_ tAccountUser) Signup(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AccountUser.Signup", args).Url
}

func (_ tAccountUser) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AccountUser.Login", args).Url
}

func (_ tAccountUser) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AccountUser.Logout", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tAdminArticle struct {}
var AdminArticle tAdminArticle


func (_ tAdminArticle) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminArticle.List", args).Url
}

func (_ tAdminArticle) Create(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("AdminArticle.Create", args).Url
}

func (_ tAdminArticle) Update(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminArticle.Update", args).Url
}

func (_ tAdminArticle) Delete(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminArticle.Delete", args).Url
}


type tAdminUser struct {}
var AdminUser tAdminUser


func (_ tAdminUser) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminUser.List", args).Url
}

func (_ tAdminUser) Create(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("AdminUser.Create", args).Url
}

func (_ tAdminUser) Update(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminUser.Update", args).Url
}

func (_ tAdminUser) Delete(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminUser.Delete", args).Url
}


