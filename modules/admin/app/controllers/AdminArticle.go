package controllers

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"gogo/app/models"
	"gogo/common/db"
	"fmt"
	"os"
	"time"
)

type AdminArticle struct {
	*revel.Controller
}

func (this AdminArticle) List() revel.Result {

	session,err := db.Connect()
	if err != nil {
		fmt.Printf("Error connection")
		os.Exit(1)
	}

	defer session.Close()

	var articles []models.Article

	err = session.DB("blog").C("articles").Find(bson.M{}).All(&articles)

	if err!=nil {
		fmt.Printf("Find error")
		os.Exit(1)
	}

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

		session,err := db.Connect()
		if err != nil {
			fmt.Printf("Error connection")
			os.Exit(1)
		}

		defer session.Close()

		err = session.DB("blog").C("articles").Insert(models.Article{
			Id: bson.NewObjectId().Hex(),
			Title: article.Title,
			Content: article.Content,
			Author: article.Author,
			Created: time.Now(),
			Updated: time.Now()})

		if err!=nil {
			fmt.Printf("Find error")
			os.Exit(1)
		}

		return this.Redirect(AdminArticle.List)
	} 

	return this.Render()
}

func (this AdminArticle) Update(id string) revel.Result {

	session,err := db.Connect()
	if err != nil {
		fmt.Printf("Error connection")
		os.Exit(1)
	}

	defer session.Close()

	if this.Request.Method == "POST" {

		var article *models.Article

		this.Params.Bind(&article, "article")

		article.Validate(this.Validation)

		if this.Validation.HasErrors() {
			this.Validation.Keep()
			this.FlashParams()
			return this.Redirect("/Article/Update/"+id)
		}		

		err = session.DB("blog").C("articles").Update(bson.M{"_id": id}, bson.M{
			"$set": bson.M{
				"title": article.Title,
				"author": article.Author,
				"content": article.Content,
				"updated": time.Now()}})

		if err != nil {
			os.Exit(1)
		}

		return this.Redirect(AdminArticle.List)
	} 	

	var article *models.Article

	err = session.DB("blog").C("articles").Find(bson.M{"_id":id}).One(&article)

	if err!=nil {
		fmt.Printf("Find error")
		os.Exit(1)
	}

	this.RenderArgs["article"] = article

	return this.Render()
}

func (this AdminArticle) Delete(id string) revel.Result {
	session,err := db.Connect()
	if err != nil {
		fmt.Printf("Error connection")
		os.Exit(1)
	}

	defer session.Close()

	err = session.DB("blog").C("articles").Remove(bson.M{"_id":id})

	if err!=nil {
		fmt.Printf("Delete error")
		os.Exit(1)
	}

	return this.Redirect(AdminArticle.List)
}
