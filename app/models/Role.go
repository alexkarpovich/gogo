package models

import (
	"github.com/revel/revel"
)

type Role struct {
	Id			string 			`bson:"_id"`
	Name 		string 			`bson:"name"`
	Permissions []Permission 	`bson:"permissions"`
}

func (this Role) Validate(v *revel.Validation) {
	v.Required(this.Name)
}

func GetPermissions() map[string]map[string][]string {
	return map[string]map[string][]string {
		"Admin": map[string][]string {
			"AdminUser": []string {
				"List",
				"Create",
				"Delete",
				"Update"},
			"AdminRole": []string {
				"List",
				"Create",
				"Delete",
				"Update"},
			"AdminArticle": []string {
				"List",
				"Create",
				"Delete",
				"Update"},
			"AccountUser": []string {
				"Signup",
				"Login",
				"Logout",
				"Profile",
				"Retrieve"},
			"App": []string {
				"Index"}},
		"Employee": map[string][]string {
			"AdminUser": []string {
				"List"},
			"AdminRole": []string {
				"List"},
			"AdminArticle": []string {
				"List",
				"Create",
				"Update"},
			"AccountUser": []string {
				"Signup",
				"Login",
				"Logout",
				"Profile",
				"Retrieve"},
			"App": []string {
				"Index"}},
		"Guest": map[string][]string {
			"AccountUser": []string {
				"Signup",
				"Login",
				"Logout",
				"Profile",
				"Retrieve"},
			"App": []string {
				"Index"}}}
}
