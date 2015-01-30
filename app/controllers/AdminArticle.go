package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"gogo/app/models"
	"time"
)

type AdminArticle struct {
	BaseController
}

func (this AdminArticle) List() revel.Result {

	var articles []models.Article

	this._FindAll("articles", &articles)

	this.RenderArgs["articles"] = articles

	return this.Render()
}

func (this AdminArticle) Create(article *models.Article) revel.Result {
	if this.Request.Method == "POST" {

		article.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect(AdminArticle.Create)
		}

		user, ok := this.RenderArgs["loggedInUser"]
		
		if !ok {
			return this.Redirect("/account/user/login")
		}

		author := user.(*models.User)

		this._Insert("articles", models.Article{
			Id: bson.NewObjectId().Hex(),
			Title: article.Title,
			Content: article.Content,
			Author: *author,
			Created: time.Now(),
			Updated: time.Now()})

		return this.Redirect(AdminArticle.List)
	} 

	return this.Render()
}

func (this AdminArticle) Update(id string) revel.Result {

	if this.Request.Method == "POST" {

		var article *models.Article

		this.Params.Bind(&article, "article")

		article.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect("/Article/Update/"+id)
		}		

		this._Update("article", bson.M{"_id": id}, bson.M{
			"$set": bson.M{
				"title": article.Title,
				"content": article.Content,
				"updated": time.Now()}})

		return this.Redirect(AdminArticle.List)
	} 	

	var article *models.Article

	this._FindOne("articles", bson.M{"_id":id}, &article)

	this.RenderArgs["article"] = article

	return this.Render()
}

func (this AdminArticle) Delete(id string) revel.Result {
	this._Delete("articles", bson.M{"_id":id})

	return this.Redirect(AdminArticle.List)
}
