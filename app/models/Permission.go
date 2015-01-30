package models

import (
	"github.com/revel/revel"
)

type Permission struct {
	Id 			string 	`bson: "_id"`
	Contoller 	string	`bson: "controller"`
	Action 		string 	`bson: "action"` 
}

func (this Permission) Validate(v *revel.Validation) {
	v.Required(this.Contoller)
	v.Required(this.Action)
}
