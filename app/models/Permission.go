package models

import (
	"github.com/revel/revel"
)

type Permission struct {
	Id       string   `bson:"_id"`
	Resource Resource `bson:"resource"`
	Roles    []Role   `bson:"roles"`
}

func (this Permission) Validate(v *revel.Validation) {
	v.Required(this.Resource)
	v.Required(this.Roles)
}
