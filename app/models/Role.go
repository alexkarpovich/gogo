package models

import (
	"github.com/revel/revel"
)

type Role struct {
	Id	string 		`bson:"_id"`
	Name string 	`bson:"name"`
}

func (this Role) Validate(v *revel.Validation) {
	v.Required(this.Name)
}
